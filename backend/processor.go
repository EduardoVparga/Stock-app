package main

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

func processTickers(ctx context.Context, pool *pgxpool.Pool) {
	rows, err := pool.Query(ctx, "SELECT ticker FROM get_active_tickers()")
	if err != nil {
		log.Printf("Error fetching active tickers: %v", err)
		return
	}
	defer rows.Close()

	var tickers []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			log.Printf("Error scanning ticker: %v", err)
			continue
		}
		tickers = append(tickers, t)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Row iteration error: %v", err)
	}

	batches := splitIntoBatches(tickers, 3)
	var wg sync.WaitGroup
	wg.Add(len(batches))

	for i, batch := range batches {
		go func(idx int, bt []string) {
			defer wg.Done()
			log.Printf("Procesando lote %d/%d: %v", idx+1, len(batches), bt)
			for _, ticker := range bt {
				price, err := FetchSpotPrice(ticker)
				if err != nil {
					log.Printf("Error fetching spot para %s: %v", ticker, err)
					continue
				}
				// UPSERT reemplaza INSERT+ON CONFLICT
				if _, err := pool.Exec(ctx, `
                    UPSERT INTO SpotPrices (ticker, price, loaded_at)
                    VALUES ($1, $2, now())
                `, ticker, price); err != nil {
					log.Printf("Error upserting spot price para %s: %v", ticker, err)
				}
			}
		}(i, batch)
	}
	wg.Wait()
}

func splitIntoBatches(items []string, num int) [][]string {
	var batches [][]string
	total := len(items)
	if total == 0 {
		return batches
	}
	batchSize := (total + num - 1) / num
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}
		batches = append(batches, items[i:end])
	}
	return batches
}
