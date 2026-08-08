package server

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type Server struct {
	cfg    *config.Config
	videos map[string]*config.Video

	mpvConn    net.Conn
	mpvEncoder *json.Encoder
	mpvMu      sync.Mutex

	playbackStatus PlaybackStatus
	clients        map[chan string]struct{}
	clientsMu      sync.Mutex

	httpServer *http.Server
}

func NewServer(cfg *config.Config) (s *Server) {
	s = &Server{
		cfg:            cfg,
		videos:         make(map[string]*config.Video, len(cfg.Videos)),
		clients:        make(map[chan string]struct{}),
		playbackStatus: PlaybackStatus{Idle: true},
	}

	s.processConfig()
	cfg.Print("Processed config")

	return
}

func (s *Server) Start() error {
	log.Println("Starting Server")

	if err := s.setupMpv(); err != nil {
		log.Printf("Error setting up mpv. Is mpv running?\n%v", err)
	}

	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("/", s.Home)
	mux.HandleFunc("/play", s.PlayHandler)
	mux.HandleFunc("GET /dialog/{id}", s.DialogHandler)
	mux.HandleFunc("/mpvStatus", s.StatusHandler)

	// Static files
	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	fs = http.FileServer(http.Dir(s.cfg.ConfigFolder))
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
