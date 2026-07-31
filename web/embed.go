package web

import "embed"

//go:embed templates/* static/*
var Files embed.FS

func Templates() embed.FS {
	return Files
}
