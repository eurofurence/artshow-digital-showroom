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

	if err := playVideo(req.Video, s.mpvEncoder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func playVideo(file string, encoder *json.Encoder) error {
	cmd := map[string]any{
		"command": []any{
			"loadfile",
			file,
			"replace",
		},
	}

	return encoder.Encode(cmd)
}
