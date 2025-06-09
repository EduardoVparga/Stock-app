// file name: backend/internal/apidata/client.go
package apidata

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"stock-app/internal/config"
	"stock-app/internal/domain"
)

// Client se comunica con la API de datos de acciones.
type Client struct {
	cfg        config.API
	httpClient *http.Client
}

// NewClient crea una nueva instancia del cliente de apidata.
func NewClient(cfg config.API) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: &http.Client{},
	}
}

type apiResponse struct {
	Items    []domain.StockInfo `json:"items"`
	NextPage string             `json:"next_page"`
}

// FetchAll recorre todas las páginas de la API y devuelve todos los ítems.
func (c *Client) FetchAll(ctx context.Context) ([]domain.StockInfo, error) {
	var all []domain.StockInfo
	nextPage := ""

	for {
		req, err := http.NewRequestWithContext(ctx, "GET", c.cfg.URL, nil)
		if err != nil {
			return nil, fmt.Errorf("creando request: %w", err)
		}
		if nextPage != "" {
			q := req.URL.Query()
			q.Add("next_page", nextPage)
			req.URL.RawQuery = q.Encode()
		}
		req.Header.Set("Authorization", "Bearer "+c.cfg.Bearer)

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("petición HTTP: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("status %s: %s", resp.Status, string(body))
		}

		var r apiResponse
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return nil, fmt.Errorf("parseando JSON: %w", err)
		}

		all = append(all, r.Items...)
		if r.NextPage == "" {
			break
		}
		nextPage = r.NextPage
	}
	return all, nil
}
