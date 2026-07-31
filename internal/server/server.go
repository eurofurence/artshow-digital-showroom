package server

import (
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/handlers"
)

type Server struct {
	config *config.Config
	addr   string
}

func New(config *config.Config) *Server {
	return &Server{config: config, addr: config.MediaInterface.Port}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Handlers
	h := handlers.New(s.config)

	// Pages
	mux.HandleFunc("/", h.Home)
	mux.HandleFunc("/hello", h.Hello)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return http.ListenAndServe(s.addr, mux)
}
