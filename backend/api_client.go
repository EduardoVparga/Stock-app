package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// spotResponse coincide con el JSON devuelto por la API.
type spotResponse struct {
	Spot float64 `json:"spot"`
}

func FetchSpotPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("http://localhost:8000/%s", ticker)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("status %s: %s", resp.Status, string(body))
	}

	var sr spotResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}
	return sr.Spot, nil
}
