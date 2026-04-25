package services

import (
	"price-comparison-backend/internal/integrations"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Platform string `json:"platform"`
	Price    string `json:"price"`
	URL      string `json:"url"`
}

func SearchProducts(query string, apiKey string, host string) ([]Product, error) {
	data, err := integrations.FetchProducts(query, apiKey, host)
	if err != nil {
		return nil, err
	}

	products := []Product{}

	for i, item := range data.Data.Products {
		if i >= 5 {
			break
		}

		products = append(products, Product{
			ID:       item.ASIN,
			Name:     item.Title,
			Platform: "Amazon",
			Price:    item.Price,
			URL:      item.URL,
		})
	}

	return products, nil
}
