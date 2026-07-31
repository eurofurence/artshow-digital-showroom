package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/web"
)

type Handler struct {
	config *config.Config
}

func New(config *config.Config) *Handler {
	return &Handler{config: config}
}

type HomePageData struct {
	Title   string
	Columns int
	Entries []config.MediaItem
}

var templates = template.Must(
	template.ParseFS(
		web.Files,
		"templates/*.html",
		"templates/partials/*.html",
	),
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("got request for", r.Method, r.URL.Path)

	entries := make([]config.MediaItem, len(h.config.MediaItems))
	copy(entries, h.config.MediaItems)
	for i := range entries {
		entries[i].Thumbnail = thumbnailFor(entries[i].Path)
	}

	data := HomePageData{
		Title:   h.config.MediaInterface.Title,
		Columns: h.config.MediaInterface.Columns,
		Entries: entries,
	}

	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "message", map[string]string{
		"Message": "Hello from Go + HTMX!",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func thumbnailFor(path string) string {
	/*
		if fileExists(path) {
			generate preview image
		}
	*/

	return "/static/fallback-lyca-shocked-bw.png"
}
