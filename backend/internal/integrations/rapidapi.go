package integrations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RapidAPIResponse struct {
	Data struct {
		Products []struct {
			ASIN  string `json:"asin"`
			Title string `json:"product_title"`
			Price string `json:"product_price"`
			URL   string `json:"product_url"`
		} `json:"products"`
	} `json:"data"`
}

func FetchProducts(query string, apiKey string, host string) (*RapidAPIResponse, error) {
	encodedQuery := url.QueryEscape(query)

	requestURL := fmt.Sprintf(
		"https://%s/search?query=%s&page=1&country=US&sort_by=RELEVANCE&product_condition=ALL&is_prime=false&deals_and_discounts=NONE",
		host,
		encodedQuery,
	)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", host)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("RapidAPI Status:", resp.StatusCode)
	fmt.Println("RapidAPI Response:", string(bodyBytes))

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("rapidapi returned status %d", resp.StatusCode)
	}

	var result RapidAPIResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
