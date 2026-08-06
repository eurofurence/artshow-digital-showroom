package server

import (
	"encoding/json"
	"log"
	"net"
)

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

	return nil
}

func (s *Server) processMpvOutput() {
	dec := json.NewDecoder(s.mpvConn)
	for {
		var msg map[string]any
		if err := dec.Decode(&msg); err != nil {
			return
		}
		log.Printf("mpv: %#v", msg)
	}
}
