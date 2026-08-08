package server

import (
	"log"
	"net/http"
)

func (s *Server) DialogHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DialogHandler got request for", r.Method, r.URL.Path)

	id := r.PathValue("id")

	video, ok := s.videos[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := s.templates.ExecuteTemplate(w, "dialog", video); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
