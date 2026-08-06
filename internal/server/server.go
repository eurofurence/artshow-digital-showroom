package server

import (
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type Server struct {
	cfg  *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", s.Home)
	mux.HandleFunc("/play", s.PlayHandler)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fs = http.FileServer(http.Dir(s.cfg.Folder))
	mux.Handle("/media/", http.StripPrefix("/media/", fs))

	return http.ListenAndServe(s.cfg.MediaInterface.Port, mux)
}
