package app

import (
	"fmt"
	db "hotel-service/config/db"
	env "hotel-service/config/env"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func NewConfig() Config {
	port := env.GetString("PORT", ":8000")
	return Config{
		Addr: port,
	}
}

func NewApplication(config Config) *Application {
	return &Application{
		Config: config,
	}
}

func Run(application *Application) error {
	_, err := db.SetupDB()
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		return fmt.Errorf("Failed to connect to Database: %w", err)
	}

	fmt.Println("Connection to DB was successful")

	server := &http.Server{
		Addr:         application.Config.Addr,
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on PORT : " + application.Config.Addr)
	return server.ListenAndServe()
}
