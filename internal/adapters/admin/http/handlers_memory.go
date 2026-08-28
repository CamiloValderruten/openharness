package adminhttp

import (
	"html/template"
	"net/http"
	"sort"
	"strings"
)

type memoryPageData struct {
	pageData
	Files        []MemoryFileItem
	SelectedPath string
	FileContent  template.HTML
	TotalCount   int
}

type MemoryFileItem struct {
	Path     string
	Filename string
	Size     int
}

func (s *Server) handleMemoryPage(w http.ResponseWriter, r *http.Request) {
	pd := s.basePageData(r, "memory")
	files := s.getMemoryFiles()
	data := memoryPageData{
		pageData:   pd,
		Files:      files,
		TotalCount: len(files),
	}
	s.render(w, "memory.html", data)
}

func (s *Server) handleFragMemoryFile(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Query().Get("path"))
	if path == "" {
		http.Error(w, "missing path", http.StatusBadRequest)
		return
	}

	content := "File not found"
	if s.deps.Memory != nil {
		if c, err := s.deps.Memory.Read(path); err == nil {
			content = c
		} else {
			content = "Error reading file: " + err.Error()
		}
	}

	data := struct {
		Path    string
		Content template.HTML
	}{
		Path:    path,
		Content: template.HTML(formatChatMarkdown(content)),
	}
	s.renderFragment(w, "frag_memory.html", data)
}

func (s *Server) getMemoryFiles() []MemoryFileItem {
	if s.deps.Memory == nil {
		return nil
	}

	all, err := s.deps.Memory.AllFiles()
	if err != nil {
		return nil
	}

	var items []MemoryFileItem
	for path, content := range all {
		// Ignore trash or hidden system files
		if strings.HasPrefix(path, ".trash") || strings.HasPrefix(path, "prompts/") {
			continue
		}
		parts := strings.Split(path, "/")
		filename := parts[len(parts)-1]
		items = append(items, MemoryFileItem{
			Path:     path,
			Filename: filename,
			Size:     len(content),
		})
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Path < items[j].Path
	})
	return items
}
