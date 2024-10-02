package app

import (
	"fmt"
	"net/http"
	"os"
	"restaurant_management/internal/routes"
	"strconv"
	"time"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	NewServer := &Server{
		port: port,
	}

	server := http.Server{
		Addr:              fmt.Sprintf(":%d", NewServer.port),
		Handler:           routes.Routes(),
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       1 * time.Second,
		MaxHeaderBytes:    8192,
	}

	return &server
}
