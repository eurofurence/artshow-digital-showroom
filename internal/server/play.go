package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type PlayRequest struct {
	Video string `json:"video"`
}

func (s *Server) PlayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for", r.Method, r.URL.Path)

	var req PlayRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Playrequest %v", req)

	var err error
	if req.Video == "" {
		err = s.stopVideo()
	} else {
		err = s.playVideo(req.Video, "my title")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) playVideo(file string, title string) error {
	return s.sendCommandsToMpv(
		[]any{"loadfile", file, "replace"},
		[]any{"show-text", title, 10_000}, // time in ms
	)
}

func (s *Server) stopVideo() error {
	return s.sendCommandsToMpv(
		[]any{"stop"}, // stop all playback
		[]any{"show-text", "", 0},
	)
}
