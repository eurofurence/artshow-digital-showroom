package server

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/web"
)

type Server struct {
	templates *template.Template

	cfg    *config.Config
	videos map[string]*config.Video

	mpvConn    net.Conn
	mpvEncoder *json.Encoder
	mpvMu      sync.Mutex

	playbackStatus PlaybackStatus
	clients        map[chan string]struct{}
	clientsMu      sync.Mutex

	httpServer *http.Server

	shutdownCtx    context.Context
	shutdownCancel context.CancelFunc
}

func NewServer(cfg *config.Config) (s *Server) {
	shutdownCtx, shutdownCancel := context.WithCancel(context.Background())

	templates := template.Must(
		template.New("").
			Funcs(template.FuncMap{"formatDuration": formatDuration}).
			ParseFS(
				web.Files,
				"templates/*.html",
				"templates/partials/*.html",
			),
	)

	s = &Server{
		templates:      templates,
		cfg:            cfg,
		videos:         make(map[string]*config.Video, len(cfg.Videos)),
		clients:        make(map[chan string]struct{}),
		playbackStatus: PlaybackStatus{Idle: true},
		shutdownCtx:    shutdownCtx,
		shutdownCancel: shutdownCancel,
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
	fs := http.FileServer(http.FS(web.Files))
	mux.Handle("/static/", fs)

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

	// Tell long-lived handlers, such as SSE, to terminate.
	s.shutdownCancel()

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if s.httpServer != nil {
		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
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
