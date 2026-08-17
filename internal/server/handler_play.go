package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/arguments"
	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/constants"
)

type PlayRequest struct {
	ID string `json:"id"`
}

func (s *Server) PlayHandler(w http.ResponseWriter, r *http.Request) {
	if arguments.Verbose() {
		log.Printf("PlayHandler got request for %s %q", r.Method, r.URL.Path)
	}

	var req PlayRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var err error
	if req.ID == "" {
		err = s.startStandby()
	} else {
		video, ok := s.videos[req.ID]
		if !ok || !video.VideoExists {
			http.Error(w, "video not found", http.StatusNotFound)
			return
		}

		err = s.playVideo(video)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) playVideo(video *config.Video) error {
	s.playbackStatus.StartTrack(video.Title)
	s.PublishStatus()

	return s.sendCommandsToMpv(
		// run once
		[]any{"set_property", "loop-file", "no"},

		[]any{"loadfile", video.Video, "replace"},
		[]any{"show-text", video.Title, constants.PlaybackTitleDuration},

		[]any{"set_property", "image-display-duration", 10},
		[]any{"loadfile", video.PostCredit, "append-play"},
	)
}

func (s *Server) startStandby() error {
	// do not set idle again
	if s.playbackStatus.IsIdle() {
		return nil
	}

	s.playbackStatus.SetIdle()
	s.PublishStatus()

	return s.sendCommandsToMpv(
		// repeat forever
		[]any{"set_property", "loop-file", "inf"},

		[]any{"loadfile", s.cfg.MediaInterface.Address + constants.StandbyVideo},
		[]any{"show-text", "", 0},
	)
}
