package server

import (
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type HomePageData struct {
	Title   string
	Columns int
	Entries []config.Video
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("HomeHandler got request for", r.Method, r.URL.Path)

	data := HomePageData{
		Title:   s.cfg.MediaInterface.Title,
		Columns: s.cfg.MediaInterface.Columns,
		Entries: s.cfg.Videos,
	}

	err := s.templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
