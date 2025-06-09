package domain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Currency es un float64 que sabe deserializarse desde una cadena tipo "$4.20".
type Currency float64

func (c *Currency) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	s = strings.ReplaceAll(s, "$", "")
	s = strings.ReplaceAll(s, ",", "")
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parsing Currency %q: %w", s, err)
	}
	*c = Currency(f)
	return nil
}

// StockInfo representa la información de una acción obtenida de la API externa.
type StockInfo struct {
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	Company    string    `json:"company"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom Currency  `json:"target_from"`
	TargetTo   Currency  `json:"target_to"`
	Ticker     string    `json:"ticker"`
	Time       time.Time `json:"time"`
}

// SpotPriceResponse coincide con el JSON devuelto por la API de precios spot.
type SpotPriceResponse struct {
	Spot float64 `json:"spot"`
}
