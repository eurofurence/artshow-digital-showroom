package server

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"strings"
)

func (s *Server) sendCommandsToMpv(cmds ...[]any) error {
	// protect against conflicting requests by using a mutex
	s.mpvMu.Lock()
	defer s.mpvMu.Unlock()

	if s.mpvEncoder == nil {
		return errors.New("mpv is not connected")
	}

	for _, cmd := range cmds {
		command := map[string]any{"command": cmd}
		if err := s.mpvEncoder.Encode(command); err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) setupMpv() error {
	conn, err := net.Dial("unix", s.cfg.Playout.MpvSocket)
	if err != nil {
		return err
	}
	s.mpvConn = conn
	s.mpvEncoder = json.NewEncoder(conn)

	go s.processMpvOutput()

	return s.sendCommandsToMpv(
		[]any{"observe_property", 1, "idle-active"},
		[]any{"observe_property", 2, "playback-time"},
		[]any{"observe_property", 3, "duration"},
		[]any{"observe_property", 4, "filename"},
	)
}

func (s *Server) processMpvOutput() {
	dec := json.NewDecoder(s.mpvConn)
	for {
		var msg map[string]any
		if err := dec.Decode(&msg); err != nil {
			log.Println("mpv connection closed:", err)
			return
		}

		if msg["event"] == "property-change" {
			// if arguments.Verbose() {
			// log.Printf("mpv: %v", msg)
			// }

			switch msg["name"] {
			case "idle-active":
				idle, ok := msg["data"].(bool)
				if !ok || !idle {
					continue
				}

				if err := s.startStandby(); err != nil {
					log.Printf("failed to start standby: %v", err)
				}
			case "playback-time":
				if s.playbackStatus.IsIdle() {
					continue
				}

				time, ok := msg["data"].(float64)
				if !ok {
					continue
				}

				seconds := int(time)
				if s.playbackStatus.SetPosition(seconds) {
					s.PublishStatus()
				}
			case "duration":
				duration, ok := msg["data"].(float64)
				if !ok {
					continue
				}

				s.playbackStatus.SetDuration(int(duration))
				s.PublishStatus()
			case "filename":
				filename, ok := msg["data"].(string)
				if !ok || !strings.HasSuffix(filename, ".png") {
					continue
				}

				s.playbackStatus.SetInCredits()
				s.PublishStatus()
			}
		}
	}
}
