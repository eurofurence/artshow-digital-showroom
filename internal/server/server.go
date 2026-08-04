package server

import (
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/handlers"
)

type Server struct {
	cfg *config.Config
	addr   string
}

func New(cfg *config.Config) *Server {
	return &Server{cfg: cfg, addr: cfg.MediaInterface.Port}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Handlers
	h := handlers.New(s.cfg)

	// Pages
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/hello", h.Hello)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fs = http.FileServer(http.Dir(s.cfg.Folder))
	mux.Handle("/media/", http.StripPrefix("/media/", fs))

	return http.ListenAndServe(s.addr, mux)
}
