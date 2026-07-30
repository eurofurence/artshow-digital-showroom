package main

import (
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/config"
	"github.com/eurofurence/artshow-digital-showroom/internal/server"
)

func main() {
	cfg, err := config.Load("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	srv := server.New(":8080")

	log.Println("Listening on http://localhost:8080")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
