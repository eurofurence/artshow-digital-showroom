package main

import (
	"log"

	"github.com/eurofurence/artshow-digital-showroom/internal/server"
)

func main() {
	srv := server.New(":8080")

	log.Println("Listening on http://localhost:8080")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
