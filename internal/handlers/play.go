package handlers

import (
	"log"
	"net"
	"net/http"
	"encoding/json"
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

func playVideo(file string) error {
	log.Println(file)

	conn, err := net.Dial("unix", "/tmp/mpv.sock")
	if err != nil {
		return err
	}
	defer conn.Close()

	cmd := map[string]any{
		"command": []any{
			"loadfile",
			file,
			"replace",
		},
	}

	return json.NewEncoder(conn).Encode(cmd)
}
