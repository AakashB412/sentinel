package main

import (
	"log"

	"github.com/AakashB412/sentinel/backend/internal/server"
)

func main() {
	srv := server.New()

	log.Println("Starting Sentinel API on :8080")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
