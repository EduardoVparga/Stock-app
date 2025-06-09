package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Currency es un float64 que sabe deserializarse desde una cadena tipo "$4.20".
type Currency float64

func (c *Currency) UnmarshalJSON(data []byte) error {
	// data viene como "\"$4.20\"" o similar
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	// Quitar signo de dólar y comas (si las hubiera)
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parsing Currency %q: %w", s, err)
	}
	*c = Currency(f)
	return nil
}

// Item ahora incluye SpotPrice, que se obtiene del intraday de Alpha Vantage.
type Item struct {
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	Company    string    `json:"company"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom Currency  `json:"target_from"`
	TargetTo   Currency  `json:"target_to"`
	Ticker     string    `json:"ticker"`
	Time       time.Time `json:"time"`

	// Nuevo campo para el precio spot más reciente:
	SpotPrice float64
}

type apiResponse struct {
	Items    []Item `json:"items"`
	NextPage string `json:"next_page"`
}

// FetchAll recorre todas las páginas de la API y devuelve todos los ítems,
// pero además, para cada Item llama a FetchSpotPrice y rellena SpotPrice.
func FetchAll(ctx context.Context, cfg APIConfig) ([]Item, error) {
	var all []Item
	next := ""

	for {
		req, err := http.NewRequestWithContext(ctx, "GET", cfg.URL, nil)
		if err != nil {
			return nil, fmt.Errorf("creando request: %w", err)
		}
		if next != "" {
			q := req.URL.Query()
			q.Add("next_page", next)
			req.URL.RawQuery = q.Encode()
		}
		req.Header.Set("Authorization", "Bearer "+cfg.Bearer)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("petición HTTP: %w", err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status %s: %s", resp.Status, string(body))
		}

		var r apiResponse
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, fmt.Errorf("parseando JSON: %w", err)
		}

		all = append(all, r.Items...)
		if r.NextPage == "" {
			break
		}
		next = r.NextPage
	}

	log.Printf("Registros obtenidos: %d\n", len(all))
	return all, nil
}
