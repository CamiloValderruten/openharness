package adminhttp

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/CamiloValderruten/openharness/internal/agent"
	"github.com/CamiloValderruten/openharness/internal/llm"
)

// ChatViewItem is a UI-friendly representation of a message turn.
type ChatViewItem struct {
	ID           string
	Role         string // "user" or "assistant" or "system"
	ThinkingRaw  string
	ThinkingHTML template.HTML
	IsThinking   bool
	Content      template.HTML
	RawText      string
	Timestamp    string
	ToolCalls    []ChatToolCallItem
}

// ChatToolCallItem is one tool call paired with its execution result.
type ChatToolCallItem struct {
	ID       string
	Name     string
	ArgsJSON string
	Result   string
	Success  bool
}

type chatPageData struct {
	pageData
	Messages    []ChatViewItem
	AgentActive bool
}

type chatMessagesData struct {
	Messages       []ChatViewItem
	AgentActive    bool
	Phase          string
	PhaseLabel     string
	PhaseClass     string
	CountdownUntil int64
}

func (s *Server) handleChatPage(w http.ResponseWriter, r *http.Request) {
	pd := s.basePageData(r, "chat")
	items, active := s.getChatItems()
	data := chatPageData{
		pageData:    pd,
		Messages:    items,
		AgentActive: active,
	}
	s.render(w, "chat.html", data)
}

func (s *Server) handleChatSend(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	msg := strings.TrimSpace(r.FormValue("message"))
	if msg == "" {
		// Read from JSON body if form is empty
		var req struct {
			Message string `json:"message"`
		}
		_ = json.NewDecoder(r.Body).Decode(&req)
		msg = strings.TrimSpace(req.Message)
	}

	if msg != "" && s.deps.Agent != nil {
		s.deps.Agent.PushUserMessage(msg)
	}

	items, active := s.getChatItems()
	var phase, label, class string
	var countdownUntil int64
	if s.deps.Agent != nil {
		phase, label, class, countdownUntil = formatPhaseInfo(s.deps.Agent.Snapshot())
	}
	data := chatMessagesData{
		Messages:       items,
		AgentActive:    active,
		Phase:          phase,
		PhaseLabel:     label,
		PhaseClass:     class,
		CountdownUntil: countdownUntil,
	}
	s.renderFragment(w, "frag_chat_messages.html", data)
}

func (s *Server) handleFragChatMessages(w http.ResponseWriter, _ *http.Request) {
	items, active := s.getChatItems()
	var phase, label, class string
	var countdownUntil int64
	if s.deps.Agent != nil {
		phase, label, class, countdownUntil = formatPhaseInfo(s.deps.Agent.Snapshot())
	}

	data := chatMessagesData{
		Messages:       items,
		AgentActive:    active,
		Phase:          phase,
		PhaseLabel:     label,
		PhaseClass:     class,
		CountdownUntil: countdownUntil,
	}
	s.renderFragment(w, "frag_chat_messages.html", data)
}

func (s *Server) handleFragChatStatus(w http.ResponseWriter, _ *http.Request) {
	var phase, label, class string
	var countdownUntil int64
	var tokens int

	if s.deps.Agent != nil {
		snap := s.deps.Agent.Snapshot()
		tokens = snap.TokenEstimate
		phase, label, class, countdownUntil = formatPhaseInfo(snap)
	} else {
		phase, label, class = "idle", "Online", "online"
	}

	data := struct {
		Phase          string
		Label          string
		Class          string
		Tokens         int
		CountdownUntil int64
	}{
		Phase:          phase,
		Label:          label,
		Class:          class,
		Tokens:         tokens,
		CountdownUntil: countdownUntil,
	}
	s.renderFragment(w, "frag_chat_status.html", data)
}

func formatPhaseInfo(snap agent.AgentSnapshot) (phase, label, class string, countdownUntil int64) {
	phase = string(snap.Phase)
	switch snap.Phase {
	case agent.PhaseGenerating:
		return "generating", "Thinking...", "generating", 0
	case agent.PhaseExecutingTool:
		if snap.CurrentTool == "sleep" && !snap.SleepUntil.IsZero() {
			rem := time.Until(snap.SleepUntil)
			if rem > 0 {
				totalSec := int(rem.Seconds())
				m := totalSec / 60
				s := totalSec % 60
				return "sleeping", fmt.Sprintf("Sleeping for %d:%02d...", m, s), "sleeping", snap.SleepUntil.Unix()
			}
			return "sleeping", "Waking up...", "sleeping", 0
		}
		if snap.CurrentTool != "" {
			return "tool", fmt.Sprintf("Running %s...", snap.CurrentTool), "tool", 0
		}
		return "tool", "Executing Tool...", "tool", 0
	case agent.PhaseCompacting:
		return "compacting", "Compacting Context...", "generating", 0
	case agent.PhaseSaving:
		return "saving", "Saving State...", "generating", 0
	case agent.PhaseStopped:
		return "stopped", "Stopped", "offline", 0
	default:
		return "idle", "Online", "online", 0
	}
}

// getChatItems aggregates raw LLM conversation messages into UI items.
func (s *Server) getChatItems() ([]ChatViewItem, bool) {
	if s.deps.Agent == nil {
		return nil, false
	}

	rawMsgs := s.deps.Agent.Messages()
	snap := s.deps.Agent.Snapshot()
	active := snap.Phase == agent.PhaseGenerating || snap.Phase == agent.PhaseExecutingTool

	if len(rawMsgs) == 0 {
		return nil, active
	}

	// Build map of tool results by tool_call_id
	toolResults := make(map[string]string)
	for _, m := range rawMsgs {
		if m.Role == llm.RoleTool && m.ToolCallID != "" {
			toolResults[m.ToolCallID] = m.Content
		}
	}

	var items []ChatViewItem
	lastWasInternal := false

	for idx, m := range rawMsgs {
		// Skip raw system prompt messages and standalone tool-role messages (tool results are attached to assistant calls)
		if m.Role == llm.RoleSystem || m.Role == llm.RoleTool {
			continue
		}

		if m.Role == llm.RoleUser {
			isInternal, cleanText := isInternalAutonomousPrompt(m.Content)
			if isInternal {
				lastWasInternal = true
				continue
			}
			lastWasInternal = false

			items = append(items, ChatViewItem{
				ID:        strings.ReplaceAll(m.Role, " ", "-") + "-" + string(rune('0'+idx%10)),
				Role:      llm.RoleUser,
				RawText:   cleanText,
				Content:   template.HTML(formatChatMarkdown(cleanText)),
				Timestamp: time.Now().Format("15:04"),
			})
			continue
		}

		if m.Role == llm.RoleAssistant {
			// If the preceding prompt was an internal loop nudge / timer, skip this background autonomous turn
			if lastWasInternal {
				continue
			}

			thinkingRaw, isThinking, cleanContent := extractThinking(m.Content)

			var formattedContent template.HTML
			if cleanContent != "" {
				formattedContent = template.HTML(formatChatMarkdown(cleanContent))
			}

			var formattedThinking template.HTML
			if thinkingRaw != "" {
				formattedThinking = template.HTML(formatChatMarkdown(thinkingRaw))
			}

			item := ChatViewItem{
				ID:           strings.ReplaceAll(m.Role, " ", "-") + "-" + string(rune('0'+idx%10)),
				Role:         m.Role,
				ThinkingRaw:  thinkingRaw,
				ThinkingHTML: formattedThinking,
				IsThinking:   isThinking,
				RawText:      cleanContent,
				Content:      formattedContent,
				Timestamp:    time.Now().Format("15:04"),
			}

			if len(m.ToolCalls) > 0 {
				for _, tc := range m.ToolCalls {
					res := toolResults[tc.ID]
					item.ToolCalls = append(item.ToolCalls, ChatToolCallItem{
						ID:       tc.ID,
						Name:     tc.Function.Name,
						ArgsJSON: tc.Function.Arguments,
						Result:   res,
						Success:  !strings.HasPrefix(res, "Error:") && !strings.HasPrefix(res, "Failed:"),
					})
				}
			}

			// Only append if there's content, thinking, or tool calls
			if cleanContent != "" || thinkingRaw != "" || isThinking || len(item.ToolCalls) > 0 {
				items = append(items, item)
			}
		}
	}

	return items, active
}

// isInternalAutonomousPrompt checks if a user turn was generated by the internal autonomous loop
// (timer nudges, continue prompts, cold-start bootstrap, background daemon alerts) rather than a human message.
func isInternalAutonomousPrompt(content string) (isInternal bool, cleanText string) {
	s := strings.TrimSpace(content)
	if strings.HasPrefix(s, "[Time:") ||
		strings.HasPrefix(s, "Cold start.") ||
		strings.HasPrefix(s, "[Daemon alert") ||
		strings.HasPrefix(s, "[Scheduled task") ||
		strings.HasPrefix(s, "[Webhook") ||
		strings.HasPrefix(s, "[Peer message") {
		return true, ""
	}

	if strings.HasPrefix(s, "[Collaborator message") {
		// Extract the actual human collaborator words
		if idx := strings.Index(s, "Your collaborator says: "); idx != -1 {
			after := s[idx+len("Your collaborator says: "):]
			if end := strings.Index(after, "\n\nReply via"); end != -1 {
				return false, strings.TrimSpace(after[:end])
			}
			return false, strings.TrimSpace(after)
		}
		return true, ""
	}

	return false, s
}

// extractThinking separates <think>...</think> reasoning blocks from user-visible response content.
func extractThinking(raw string) (thinking string, isThinking bool, content string) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false, ""
	}

	lower := strings.ToLower(raw)
	startIdx := strings.Index(lower, "<think>")
	endIdx := strings.Index(lower, "</think>")

	if startIdx != -1 && endIdx != -1 && endIdx > startIdx {
		// Standard <think>...</think>
		before := strings.TrimSpace(raw[:startIdx])
		thinking = strings.TrimSpace(raw[startIdx+7 : endIdx])
		after := strings.TrimSpace(raw[endIdx+8:])
		if before != "" && after != "" {
			content = before + "\n\n" + after
		} else if before != "" {
			content = before
		} else {
			content = after
		}
	} else if startIdx != -1 && endIdx == -1 {
		// Currently thinking in progress (<think> without </think>)
		before := strings.TrimSpace(raw[:startIdx])
		thinking = strings.TrimSpace(raw[startIdx+7:])
		isThinking = true
		content = before
	} else if startIdx == -1 && endIdx != -1 {
		// Implicit opening <think>, closing </think> present
		thinking = strings.TrimSpace(raw[:endIdx])
		content = strings.TrimSpace(raw[endIdx+8:])
	} else {
		// No think tags
		content = raw
	}

	// Clean any stray tags
	content = strings.ReplaceAll(content, "<think>", "")
	content = strings.ReplaceAll(content, "</think>", "")
	content = strings.ReplaceAll(content, "<THINK>", "")
	content = strings.ReplaceAll(content, "</THINK>", "")
	content = strings.TrimSpace(content)

	thinking = strings.ReplaceAll(thinking, "<think>", "")
	thinking = strings.ReplaceAll(thinking, "</think>", "")
	thinking = strings.ReplaceAll(thinking, "<THINK>", "")
	thinking = strings.ReplaceAll(thinking, "</THINK>", "")
	thinking = strings.TrimSpace(thinking)

	return thinking, isThinking, content
}

// formatChatMarkdown converts basic markdown syntax (code blocks, bold, newlines, etc.) to safe HTML.
func formatChatMarkdown(text string) string {
	if text == "" {
		return ""
	}

	// Escape basic HTML
	escaped := template.HTMLEscapeString(text)

	// Replace fenced code blocks
	var result strings.Builder
	lines := strings.Split(escaped, "\n")
	inCodeBlock := false

	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			if !inCodeBlock {
				inCodeBlock = true
				result.WriteString("<pre><code>")
			} else {
				inCodeBlock = false
				result.WriteString("</code></pre>")
			}
			continue
		}

		if inCodeBlock {
			result.WriteString(line + "\n")
		} else {
			// Inline code `...`
			formattedLine := replaceInlineCode(line)
			// Bold **...**
			formattedLine = replaceBold(formattedLine)
			if formattedLine == "" {
				result.WriteString("<p></p>")
			} else {
				result.WriteString("<p>" + formattedLine + "</p>")
			}
		}
	}

	if inCodeBlock {
		result.WriteString("</code></pre>")
	}

	return result.String()
}

func replaceInlineCode(s string) string {
	parts := strings.Split(s, "`")
	if len(parts) < 3 {
		return s
	}
	var res strings.Builder
	for i, part := range parts {
		if i%2 == 1 {
			res.WriteString("<code>" + part + "</code>")
		} else {
			res.WriteString(part)
		}
	}
	return res.String()
}

func replaceBold(s string) string {
	parts := strings.Split(s, "**")
	if len(parts) < 3 {
		return s
	}
	var res strings.Builder
	for i, part := range parts {
		if i%2 == 1 {
			res.WriteString("<strong>" + part + "</strong>")
		} else {
			res.WriteString(part)
		}
	}
	return res.String()
}
