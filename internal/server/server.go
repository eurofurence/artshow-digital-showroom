package server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type Server struct {
	cfg        *config.Config
	mpvConn    net.Conn
	mpvEncoder *json.Encoder
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	log.Println("Starting Server")

	conn, err := net.Dial("unix", s.cfg.Playout.MpvSocket)
	if err != nil {
		return err
	}
	s.mpvConn = conn
	s.mpvEncoder = json.NewEncoder(conn)

	go func() {
		dec := json.NewDecoder(s.mpvConn)
		for {
			var msg map[string]any
			if err := dec.Decode(&msg); err != nil {
				return
			}
			log.Printf("mpv: %#v", msg)
		}
	}()

	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", s.Home)
	mux.HandleFunc("/play", s.PlayHandler)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fs = http.FileServer(http.Dir(s.cfg.Folder))
	mux.Handle("/media/", http.StripPrefix("/media/", fs))

	s.httpServer = &http.Server{
		Addr:    s.cfg.MediaInterface.Port,
		Handler: mux,
	}
	log.Println("Listening on " + s.cfg.MediaInterface.Address)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Close() error {
	log.Println("Shutting down")

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("http server result %v", err)
		}
	}

	if s.mpvConn != nil {
		if err := s.mpvConn.Close(); err != nil {
			log.Printf("mpv connection result %v", err)
		}
	}

	return nil
}
