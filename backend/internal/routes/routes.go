package routes

import (
	"net/http"

	"price-comparison-backend/internal/handlers"
)

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/health", handlers.HealthCheck)

	return mux
}
