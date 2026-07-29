package server

import (
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/handlers"
)

type Server struct {
	addr string
}

func New(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Pages
	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/hello", handlers.Hello)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return http.ListenAndServe(s.addr, mux)
}
