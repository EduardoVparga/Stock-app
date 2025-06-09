package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func startScheduler(ctx context.Context, pool *pgxpool.Pool) {
	ticker := time.NewTicker(30 * time.Minute)
	defer ticker.Stop()

	go processTickers(ctx, pool)

	for {
		select {
		case <-ticker.C:
			go processTickers(ctx, pool)
		case <-ctx.Done():
			log.Println("Scheduler detenido")
			return
		}
	}
}
