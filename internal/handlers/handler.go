package handlers

import (
	"github.com/eurofurence/artshow-digital-showroom/internal/config"
)

type Handler struct {
	cfg *config.Config
}

func New(config *config.Config) *Handler {
	return &Handler{cfg: config}
}
