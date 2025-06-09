package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	apiCfg := LoadAPIConfig()
	dbCfg := LoadDBConfig()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbCfg.Name, dbCfg.SSLMode,
	)
	ctx := context.Background()
	// Usar pgxpool para manejar conexiones concurrentes
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la BD: %v", err)
	}
	defer pool.Close()

	var lastLoad time.Time
	err = pool.QueryRow(ctx, `SELECT LastLoadTime()`).Scan(&lastLoad)
	if err != nil {
		log.Fatalf("error obteniendo última carga: %v", err)
	}

	go watchForUpdates(ctx, pool, lastLoad)

	if !lastLoad.IsZero() && time.Since(lastLoad) < time.Hour {
		log.Println("Última carga hace menos de una hora. Saltando actualización.")
		ServidorHTTP(ctx, pool)
		return
	}

	if _, err := pool.Exec(ctx, `CALL ArchiveAndTruncate()`); err != nil {
		log.Fatalf("error en ArchiveAndTruncate: %v", err)
	}

	items, err := FetchAll(ctx, apiCfg)
	if err != nil {
		log.Fatalf("Error al extraer datos: %v", err)
	}

	loadedAt := time.Now()
	if err := InsertItems(ctx, pool, items, loadedAt); err != nil {
		log.Fatalf("Error insertando datos: %v", err)
	}

	if _, err := pool.Exec(ctx, `CALL sync_symbols()`); err != nil {
		log.Fatalf("error en sync_symbols: %v", err)
	}

	log.Println("✅ Datos actualizados y almacenados en StockInfo")

	ServidorHTTP(ctx, pool)
}
