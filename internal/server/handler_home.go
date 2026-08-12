package server

import (
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type HomePageData struct {
	Entries []config.Video
}

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("HomeHandler got request for %s %q", r.Method, r.URL.Path)

	data := HomePageData{
		Entries: s.cfg.Videos,
	}

	err := s.templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
