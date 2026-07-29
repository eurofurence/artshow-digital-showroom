package web

import "embed"

//go:embed templates/*.html templates/partials/*.html static/*
var Files embed.FS

func Templates() embed.FS {
	return Files
}
