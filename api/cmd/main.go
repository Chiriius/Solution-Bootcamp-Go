package main

import (
	"bootcamp_api/api/server"
	"log"
)

func main() {
	server, err := server.New(":8080", ":50051")
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer server.Close()

	if err := server.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
