package adminhttp

import (
	"net/http"

	"github.com/CamiloValderruten/openharness/internal/tools"
)

type toolsPageData struct {
	pageData
	Events []tools.ToolCallEvent
}

func (s *Server) handleToolsPage(w http.ResponseWriter, r *http.Request) {
	pd := s.basePageData(r, "tools")
	var events []tools.ToolCallEvent
	if s.deps.Tools != nil {
		events = reverseEvents(s.deps.Tools.SnapshotRecent(100))
	}
	data := toolsPageData{
		pageData: pd,
		Events:   events,
	}
	s.render(w, "tools.html", data)
}
