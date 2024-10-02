package app

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/http"
	"restaurant_management/internal/config/database"
	"restaurant_management/internal/routes"
	"strconv"
	"time"
)

type Server struct {
	port int
	db   *gorm.DB
}

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func NewServer() *http.Server {
	port, err := strconv.Atoi(viper.GetString("PORT"))
	if err != nil {
		port = 8080
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Failed to conect to database: %v", err)
	}

	NewServer := &Server{
		port: port,
		db:   db,
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
