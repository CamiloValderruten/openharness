package adminhttp

import (
	"net/http"
	"runtime"
	"time"

	"github.com/CamiloValderruten/openharness/internal/agent"
	"github.com/CamiloValderruten/openharness/internal/version"
)

// pageData is the common set of fields every full-page render needs:
// auth state for the layout, navigation metadata, and status info.
type pageData struct {
	Title         string
	Authenticated bool
	Username      string
	CSRFToken     string

	// TopTab selects the top-level mode: "chat" or "admin".
	TopTab string

	// Section is the admin sidebar / sub-tab section's key.
	Section string
	// SectionLabel is the human-facing title.
	SectionLabel string

	// Nav is the list of admin section items.
	Nav []navItem

	// Common stats surfaced in the header / footer.
	Version      string
	GoVersion    string
	Uptime       string
	StartedAt    string
	StatusPhase  string
	StatusLabel  string
	StatusClass  string
	TokenEst     int
	MaxTokens    int
	TotalChats   int64
	TotalTools   int64
}

type navItem struct {
	Key    string
	Href   string
	Label  string
	Icon   string
	Active bool
}

// navItems is the canonical admin navigation list.
var navItems = []navItem{
	{Key: "dashboard", Href: "/admin", Label: "Dashboard", Icon: "📊"},
	{Key: "memory", Href: "/admin/memory", Label: "Memory Explorer", Icon: "🧠"},
	{Key: "skills", Href: "/admin/skills", Label: "Skills Catalog", Icon: "🧩"},
	{Key: "subagents", Href: "/admin/subagents", Label: "Subagents Hub", Icon: "🤖"},
	{Key: "tools", Href: "/admin/tools", Label: "Tool Trace", Icon: "📜"},
	{Key: "configuration", Href: "/admin/configuration", Label: "Configuration", Icon: "⚙️"},
	{Key: "version", Href: "/admin/version", Label: "Version & Updates", Icon: "🔄"},
	{Key: "logs", Href: "/admin/logs", Label: "System Logs", Icon: "📋"},
}

// sectionLabels maps section keys to human readable titles.
var sectionLabels = map[string]string{
	"chat":          "Chat with OpenHarness",
	"dashboard":     "System Overview",
	"memory":        "Memory Explorer",
	"configuration": "Configuration Editor",
	"subagents":     "Subagents Hub",
	"skills":        "Skills Catalog",
	"tools":         "Tool Execution Trace",
	"version":       "Version & Self-Update",
	"logs":          "Live System Logs",
}

// basePageData fills the boilerplate every authenticated page needs.
func (s *Server) basePageData(r *http.Request, section string) pageData {
	sess := sessionFromContext(r.Context())
	uptime := time.Since(s.deps.StartedAt).Round(time.Second)

	nav := make([]navItem, len(navItems))
	for i, it := range navItems {
		it.Active = it.Key == section
		nav[i] = it
	}

	label, ok := sectionLabels[section]
	if !ok {
		label = section
	}

	topTab := "admin"
	if section == "chat" {
		topTab = "chat"
	}

	statusPhase := "idle"
	statusLabel := "Online (Idle)"
	statusClass := "online"
	var tokenEst, maxTokens int
	var totalChats, totalTools int64

	if s.deps.Agent != nil {
		snap := s.deps.Agent.Snapshot()
		tokenEst = snap.TokenEstimate
		maxTokens = snap.MaxTokens
		totalChats = snap.TotalChats
		totalTools = snap.TotalToolCalls
		switch snap.Phase {
		case agent.PhaseGenerating:
			statusPhase = "generating"
			statusLabel = "Thinking..."
			statusClass = "generating"
		case agent.PhaseExecutingTool:
			statusPhase = "tool"
			statusLabel = "Executing Tool..."
			statusClass = "tool"
		case agent.PhaseCompacting:
			statusPhase = "compacting"
			statusLabel = "Compacting Context..."
			statusClass = "generating"
		case agent.PhaseStopped:
			statusPhase = "stopped"
			statusLabel = "Stopped"
			statusClass = "offline"
		default:
			statusPhase = "idle"
			statusLabel = "Online"
			statusClass = "online"
		}
	}

	pd := pageData{
		Title:         label,
		Authenticated: true,
		TopTab:        topTab,
		Section:       section,
		SectionLabel:  label,
		Nav:           nav,
		Version:       version.String(),
		GoVersion:     runtime.Version(),
		Uptime:        uptime.String(),
		StartedAt:     s.deps.StartedAt.UTC().Format(time.RFC3339),
		StatusPhase:   statusPhase,
		StatusLabel:   statusLabel,
		StatusClass:   statusClass,
		TokenEst:      tokenEst,
		MaxTokens:     maxTokens,
		TotalChats:    totalChats,
		TotalTools:    totalTools,
	}
	if sess != nil {
		pd.Username = sess.Username
		pd.CSRFToken = sess.CSRFToken
	}
	return pd
}
