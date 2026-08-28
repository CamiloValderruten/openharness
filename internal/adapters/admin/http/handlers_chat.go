package adminhttp

import (
	"encoding/json"
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
	Messages    []ChatViewItem
	AgentActive bool
	Phase       string
	PhaseLabel  string
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
		http.Error(w, "invalid form", http.StatusBadRequest)
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
	data := chatMessagesData{
		Messages:    items,
		AgentActive: active,
	}
	s.renderFragment(w, "frag_chat_messages.html", data)
}

func (s *Server) handleFragChatMessages(w http.ResponseWriter, _ *http.Request) {
	items, active := s.getChatItems()
	phase := "idle"
	label := "Online"
	if s.deps.Agent != nil {
		snap := s.deps.Agent.Snapshot()
		phase = string(snap.Phase)
		switch snap.Phase {
		case agent.PhaseGenerating:
			label = "Thinking..."
		case agent.PhaseExecutingTool:
			label = "Executing Tool..."
		case agent.PhaseCompacting:
			label = "Compacting..."
		}
	}

	data := chatMessagesData{
		Messages:    items,
		AgentActive: active,
		Phase:       phase,
		PhaseLabel:  label,
	}
	s.renderFragment(w, "frag_chat_messages.html", data)
}

func (s *Server) handleFragChatStatus(w http.ResponseWriter, _ *http.Request) {
	phase := "idle"
	label := "Online (Idle)"
	class := "online"
	var tokens int

	if s.deps.Agent != nil {
		snap := s.deps.Agent.Snapshot()
		tokens = snap.TokenEstimate
		switch snap.Phase {
		case agent.PhaseGenerating:
			phase = "generating"
			label = "Thinking..."
			class = "generating"
		case agent.PhaseExecutingTool:
			phase = "tool"
			label = "Executing Tool..."
			class = "tool"
		case agent.PhaseCompacting:
			phase = "compacting"
			label = "Compacting Context..."
			class = "generating"
		default:
			phase = "idle"
			label = "Online"
			class = "online"
		}
	}

	data := struct {
		Phase  string
		Label  string
		Class  string
		Tokens int
	}{
		Phase:  phase,
		Label:  label,
		Class:  class,
		Tokens: tokens,
	}
	s.renderFragment(w, "frag_chat_status.html", data)
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
	for idx, m := range rawMsgs {
		// Skip raw system prompt messages and standalone tool-role messages (tool results are attached to assistant calls)
		if m.Role == llm.RoleSystem || m.Role == llm.RoleTool {
			continue
		}

		// Skip internal automated continue/nudge prompts that have no user message
		if m.Role == llm.RoleUser && (strings.HasPrefix(m.Content, "[Time:") && strings.Contains(m.Content, "continue")) {
			continue
		}

		thinkingRaw, isThinking, cleanContent := extractThinking(m.Content)

		item := ChatViewItem{
			ID:           strings.ReplaceAll(m.Role, " ", "-") + "-" + strings.TrimSpace(strings.ReplaceAll(m.Content, "\n", ""))[:min(8, len(strings.TrimSpace(strings.ReplaceAll(m.Content, "\n", ""))))] + "-" + string(rune('0'+idx%10)),
			Role:         m.Role,
			ThinkingRaw:  thinkingRaw,
			ThinkingHTML: template.HTML(formatChatMarkdown(thinkingRaw)),
			IsThinking:   isThinking,
			RawText:      cleanContent,
			Content:      template.HTML(formatChatMarkdown(cleanContent)),
			Timestamp:    time.Now().Format("15:04"),
		}

		if m.Role == llm.RoleAssistant && len(m.ToolCalls) > 0 {
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
		if strings.TrimSpace(cleanContent) != "" || strings.TrimSpace(thinkingRaw) != "" || isThinking || len(item.ToolCalls) > 0 {
			items = append(items, item)
		}
	}

	return items, active
}

// extractThinking separates <think>...</think> reasoning blocks from user-visible response content.
func extractThinking(raw string) (thinking string, isThinking bool, content string) {
	lower := strings.ToLower(raw)
	startIdx := strings.Index(lower, "<think>")
	if startIdx == -1 {
		return "", false, strings.TrimSpace(raw)
	}

	before := strings.TrimSpace(raw[:startIdx])
	afterStart := raw[startIdx+7:]
	lowerAfter := strings.ToLower(afterStart)

	endIdx := strings.Index(lowerAfter, "</think>")
	if endIdx == -1 {
		// Currently in-progress thinking block
		thinking = strings.TrimSpace(afterStart)
		return thinking, true, before
	}

	thinking = strings.TrimSpace(afterStart[:endIdx])
	afterEnd := strings.TrimSpace(afterStart[endIdx+8:])
	if before != "" && afterEnd != "" {
		content = before + "\n\n" + afterEnd
	} else if before != "" {
		content = before
	} else {
		content = afterEnd
	}

	return thinking, false, content
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
