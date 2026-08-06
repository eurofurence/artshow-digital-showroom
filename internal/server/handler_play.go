package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
)

type PlayRequest struct {
	Video      string `json:"video"`
	Title      string `json:"title"`
	PostCredit string `json:"postcredit"`
}

func (s *Server) PlayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for", r.Method, r.URL.Path)

	var req PlayRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Play Request %v", req)

	var err error
	if req.Video == "" {
		err = s.startStandby()
	} else {
		err = s.playVideo(req)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) playVideo(req PlayRequest) error {
	log.Printf("play %#v", req)
	return s.sendCommandsToMpv(
		// run once
		[]any{"set_property", "loop-file", "no"},

		[]any{"loadfile", req.Video, "replace"},
		[]any{"show-text", req.Title, constants.PlaybackTitleDuration},

		[]any{"set_property", "image-display-duration", 10},
		[]any{"loadfile", req.PostCredit, "append-play"},
	)
}

func (s *Server) startStandby() error {
	return s.sendCommandsToMpv(
		// repeat forever
		[]any{"set_property", "loop-file", "inf"},

		[]any{"loadfile", s.cfg.MediaInterface.Address + constants.StandbyVideo},
		[]any{"show-text", "", 0},
	)
}
