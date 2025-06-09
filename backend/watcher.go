package main

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func watchForUpdates(ctx context.Context, pool *pgxpool.Pool, lastLoad time.Time) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			var newLoad time.Time
			if err := pool.QueryRow(ctx, `SELECT LastLoadTime()`).Scan(&newLoad); err != nil {
				log.Printf("Error revisando última carga: %v", err)
				continue
			}
			if newLoad.After(lastLoad) {
				log.Println("Nueva carga detectada. Iniciando scheduler...")
				startScheduler(ctx, pool)
				return
			}
		case <-ctx.Done():
			log.Println("Watcher detenido")
			return
		}
	}
}
