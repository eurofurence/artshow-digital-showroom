package server

import (
	"encoding/json"
	"log"
	"net"
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

	route := s.cfg.MediaInterface.Address + "/" + req.Video
	if err := playVideo(route, s.cfg.Playout.MpvSocket); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func playVideo(file string, socket string) (err error) {
	log.Println(file)

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	cmd := map[string]any{
		"command": []any{
			"loadfile",
			file,
			"replace",
		},
	}

	err = json.NewEncoder(conn).Encode(cmd)
	if err != nil {
		return err
	}

	var resp map[string]any
	if err := json.NewDecoder(conn).Decode(&resp); err != nil {
		return err
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	log.Println(string(s))
	// log.Printf("response: %#v", resp)
	return
}
