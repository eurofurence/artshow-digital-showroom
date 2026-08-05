package main

import (
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/server"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New(cfg)

	log.Println("Listening on " + cfg.MediaInterface.Address)

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
