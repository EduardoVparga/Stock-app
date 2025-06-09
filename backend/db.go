package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func InsertItems(ctx context.Context, pool *pgxpool.Pool, items []Item, loadedAt time.Time) error {
	for _, it := range items {
		_, err := pool.Exec(ctx, `
        INSERT INTO StockInfo
          (action, brokerage, company, rating_from, rating_to,
           target_from, target_to, ticker, time, loaded_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
        `,
			it.Action,
			it.Brokerage,
			it.Company,
			it.RatingFrom,
			it.RatingTo,
			float64(it.TargetFrom),
			float64(it.TargetTo),
			it.Ticker,
			it.Time,
			loadedAt,
		)
		if err != nil {
			return fmt.Errorf("al insertar %v: %w", it, err)
		}
	}
	return nil
}
