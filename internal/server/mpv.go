package server

import (
	"encoding/json"
	"log"
	"net"
)

const observerID = 1

func (s *Server) sendCommandsToMpv(cmds ...[]any) error {
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
		[]any{"observe_property", observerID, "idle-active"},
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

		log.Printf("mpv: %#v", msg)

		if msg["event"] == "property-change" && msg["name"] == "idle-active" {
			idle, ok := msg["data"].(bool)
			if !ok || !idle {
				continue
			}

			log.Println("playlist finished, starting standby")

			if err := s.startStandby(); err != nil {
				log.Printf("failed to start standby: %v", err)
			}
		}
	}
}
