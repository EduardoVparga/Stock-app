package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
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
func FetchAll(ctx context.Context, cfg APIConfig, avAPIKey string) ([]Item, error) {
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

		// Por cada ítem, obtener el spot price
		for i := range r.Items {
			sp, err := FetchSpotPrice(r.Items[i].Ticker, avAPIKey)
			if err != nil {
				// Si hay error, dejamos SpotPrice en 0 y seguimos
				log.Printf("warning: no se pudo obtener spot price para %s: %v\n", r.Items[i].Ticker, err)
				r.Items[i].SpotPrice = 0
			} else {
				r.Items[i].SpotPrice = sp
			}
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

// FetchSpotPrice llama al endpoint de Intraday de Alpha Vantage y devuelve
// el “close” más reciente (último timestamp) como float64.
// endpoint: https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=<TICKER>&interval=5min&apikey=<API-KEY>
func FetchSpotPrice(ticker, apiKey string) (float64, error) {
	// Construir URL
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&outputsize=compact&apikey=%s",
		ticker, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error en GET AlphaVantage: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("status %s: %s", resp.Status, string(b))
	}

	body, _ := io.ReadAll(resp.Body)
	// Estructura parcial para parsear únicamente “Time Series (5min)”
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		return 0, fmt.Errorf("no se pudo leer JSON intraday: %w", err)
	}

	// La clave exacta es "Time Series (5min)"
	var tsField json.RawMessage
	if val, ok := raw["Time Series (5min)"]; ok {
		tsField = val
	} else {
		return 0, fmt.Errorf("respuesta no contiene Time Series (5min)")
	}

	// Deserializar cada timestamp → { "1. open": "...", "2. high": "...", ... }
	var tsMap map[string]map[string]string
	if err := json.Unmarshal(tsField, &tsMap); err != nil {
		return 0, fmt.Errorf("error parseando Time Series (5min): %w", err)
	}

	// Los keys de tsMap son strings como "2025-06-05 19:20:00".
	// Tomamos el timestamp mayor (más reciente).
	timestamps := make([]string, 0, len(tsMap))
	for k := range tsMap {
		timestamps = append(timestamps, k)
	}
	sort.Strings(timestamps) // orden lexicográfico equivale a orden cronológico aquí

	if len(timestamps) == 0 {
		return 0, fmt.Errorf("sin datos en Time Series (5min)")
	}

	last := timestamps[len(timestamps)-1] // e.g. "2025-06-05 19:20:00"
	closeStr := tsMap[last]["4. close"]   // e.g. "1.0000"
	closeVal, err := strconv.ParseFloat(closeStr, 64)
	if err != nil {
		return 0, fmt.Errorf("no se pudo convertir close %q: %w", closeStr, err)
	}

	return closeVal, nil
}
