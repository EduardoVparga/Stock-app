package spotclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"stock-app/internal/domain"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) FetchSpotPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, ticker)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return 0, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("status %s: %s", resp.Status, string(body))
	}

	var sr domain.SpotPriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&sr); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}
	return sr.Spot, nil
}
