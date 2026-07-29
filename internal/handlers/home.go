package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/eurofurence/artshow-digital-showroom/web"
)

var templates = template.Must(
	template.ParseFS(
		web.Files,
		"templates/*.html",
		"templates/partials/*.html",
	),
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("HOME:", r.Method, r.URL.Path)
	err := templates.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "message", map[string]string{
		"Message": "Hello from Go + HTMX!",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
