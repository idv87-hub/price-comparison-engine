package handlers

import (
	"encoding/json"
	"net/http"

	"price-comparison-backend/internal/config"
	"price-comparison-backend/internal/services"
)

type ProductSearchResponse struct {
	Query   string             `json:"query"`
	Results []services.Product `json:"results"`
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query().Get("query")

	if query == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "query parameter is required",
		})
		return
	}

	cfg := config.LoadConfig()

	results, err := services.SearchProducts(query, cfg.RapidAPIKey, cfg.RapidAPIHost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to fetch products",
		})
		return
	}

	response := ProductSearchResponse{
		Query:   query,
		Results: results,
	}

	json.NewEncoder(w).Encode(response)
}
