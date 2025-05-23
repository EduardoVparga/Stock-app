package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	apiCfg := LoadAPIConfig()
	dbCfg := LoadDBConfig()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbCfg.User, dbCfg.Pass, dbCfg.Host, dbCfg.Port, dbCfg.Name, dbCfg.SSLMode,
	)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la BD: %v", err)
	}
	defer conn.Close(ctx)

	if err := EnsureTable(ctx, conn); err != nil {
		log.Fatalf("Error creando tabla StockInfo: %v", err)
	}
	if err := EnsureHistoryTable(ctx, conn); err != nil {
		log.Fatalf("Error creando tabla de historial: %v", err)
	}

	// Verificar antes de consultar API
	lastLoad, err := LastLoadTime(ctx, conn)
	if err != nil {
		log.Printf("Error obteniendo última carga: %v", err)
	}
	if !lastLoad.IsZero() && time.Since(lastLoad) < time.Hour {
		log.Println("Última carga hace menos de una hora. Saltando actualización.")
		ServidorHTTP(ctx, conn)
		return
	}

	// Archivar y limpiar solo si corresponde
	if err := ArchiveCurrent(ctx, conn); err != nil {
		log.Fatalf("Error archivando datos: %v", err)
	}
	if err := TruncateCurrent(ctx, conn); err != nil {
		log.Fatalf("Error truncando tabla StockInfo: %v", err)
	}

	// Solo ahora consultar todas las páginas de la API
	items, err := FetchAll(ctx, apiCfg)
	if err != nil {
		log.Fatalf("Error al extraer datos: %v", err)
	}

	loadedAt := time.Now()
	if err := InsertItems(ctx, conn, items, loadedAt); err != nil {
		log.Fatalf("Error insertando datos: %v", err)
	}

	log.Println("✅ Datos actualizados y almacenados en StockInfo")

	ServidorHTTP(ctx, conn)
}
