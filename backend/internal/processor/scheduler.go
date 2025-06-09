// file name: backend/internal/processor/scheduler.go
package processor

import (
	"context"
	"log"
	"sync"
	"time"

	"stock-app/internal/spotclient"
	"stock-app/internal/storage"
)

// Scheduler gestiona la actualización periódica de los precios spot.
type Scheduler struct {
	repo       *storage.Repository
	spotClient *spotclient.Client
}

// NewScheduler crea una nueva instancia del planificador.
func NewScheduler(repo *storage.Repository, spotClient *spotclient.Client) *Scheduler {
	return &Scheduler{
		repo:       repo,
		spotClient: spotClient,
	}
}

// Start inicia el ciclo de actualización periódica.
func (s *Scheduler) Start(ctx context.Context) {
	log.Println("▶️  Iniciando el planificador de actualización de precios...")
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	// Ejecutar una vez al inicio
	go s.processTickers(ctx)

	for {
		select {
		case <-ticker.C:
			go s.processTickers(ctx)
		case <-ctx.Done():
			log.Println("Scheduler detenido")
			return
		}
	}
}

func (s *Scheduler) processTickers(ctx context.Context) {
	log.Println("🔄 Procesando tickers para actualizar precios spot...")
	tickers, err := s.repo.GetActiveTickers(ctx)
	if err != nil {
		log.Printf("Error fetching active tickers: %v", err)
		return
	}

	if len(tickers) == 0 {
		log.Println("No hay tickers activos para procesar.")
		return
	}

	// Procesar en lotes concurrentes para no sobrecargar la API de precios
	batches := splitIntoBatches(tickers, 3)
	var wg sync.WaitGroup
	wg.Add(len(batches))

	for i, batch := range batches {
		go func(idx int, bt []string) {
			defer wg.Done()
			log.Printf("Procesando lote %d/%d con %d tickers", idx+1, len(batches), len(bt))
			for _, ticker := range bt {
				price, err := s.spotClient.FetchSpotPrice(ticker)
				if err != nil {
					log.Printf("Error fetching spot para %s: %v", ticker, err)
					continue
				}
				if err := s.repo.UpsertSpotPrice(ctx, ticker, price); err != nil {
					log.Printf("Error upserting spot price para %s: %v", ticker, err)
				}
			}
		}(i, batch)
	}
	wg.Wait()
	log.Println("✅ Tickers procesados.")
}

func splitIntoBatches(items []string, numBatches int) [][]string {
	if len(items) == 0 || numBatches <= 0 {
		return nil
	}
	var batches [][]string
	batchSize := (len(items) + numBatches - 1) / numBatches
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}
		batches = append(batches, items[i:end])
	}
	return batches
}
