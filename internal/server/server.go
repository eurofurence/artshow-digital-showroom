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
	"github.com/eurofurence/artshow-digital-showroom/internal/stats"
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

	logger *stats.Logger

	httpServer *http.Server

	serverCtx    context.Context
	serverCancel context.CancelFunc
}

func NewServer(cfg *config.Config, ctx context.Context) (s *Server) {
	serverCtx, serverCancel := context.WithCancel(ctx)

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
		playbackStatus: PlaybackStatus{Idle: true},
		clients:        make(map[chan string]struct{}),
		serverCtx:      serverCtx,
		serverCancel:   serverCancel,
	}

	s.processConfig()
	cfg.Print("Processed config")

	return
}

func (s *Server) Start() error {
	log.Println("Starting Server")

	logger, err := stats.Open("log.csv")
	if err != nil {
		return err
	}
	s.logger = logger

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
	prefix := "/" + s.cfg.ConfigFolder + "/"
	mux.Handle(prefix, http.StripPrefix(prefix, fs))

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
	s.serverCancel()

	if err := s.logger.Close(); err != nil {
		log.Printf("csv writer failed to close %v", err)
	}

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
