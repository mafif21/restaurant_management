package main

import (
	"log"
	"restaurant_management/internal/app"
)

func main() {
	server := app.NewServer()
	log.Println("Starting server...", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
