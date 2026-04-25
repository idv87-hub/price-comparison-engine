package main

import (
	"log"
	"net/http"

	"price-comparison-backend/internal/config"
	"price-comparison-backend/internal/routes"
)

func main() {
	cfg := config.LoadConfig()

	router := routes.SetupRoutes()

	log.Printf("Starting %s in %s mode on port %s", cfg.AppName, cfg.Env, cfg.Port)

	err := http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}