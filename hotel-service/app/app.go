package app

import (
	"fmt"
	db "hotel-service/config/db"
	env "hotel-service/config/env"
	"hotel-service/controller"
	"hotel-service/repository"
	router "hotel-service/routes"
	"hotel-service/service"
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
	db, err := db.SetupDB()

	hr := repository.NewHotelRepositoryImpl(db)
	hs := service.NewHotelService(&hr)
	hc := controller.NewHotelController(hs)
	hotelRouter := router.NewHotelRouter(&hc)
	router := router.SetupRouter(&hotelRouter)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		return fmt.Errorf("Failed to connect to Database: %w", err)
	}

	fmt.Println("Connection to DB was successful")

	server := &http.Server{
		Addr:         application.Config.Addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on PORT : " + application.Config.Addr)
	return server.ListenAndServe()
}
