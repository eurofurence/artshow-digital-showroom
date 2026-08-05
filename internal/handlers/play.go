package handlers

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
)

type PlayRequest struct {
	Video string `json:"video"`
}

func (h *Handler) PlayHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for", r.Method, r.URL.Path)

	var req PlayRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	route := h.cfg.MediaInterface.Address + "/" + req.Video
	if err := playVideo(route); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func playVideo(file string) (err error) {
	log.Println(file)

	conn, err := net.Dial("unix", "/tmp/mpv.sock")
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
	return
}
