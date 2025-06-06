package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

// EnsureTable crea la tabla principal con spot_price.
func EnsureTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS StockInfo (
      action       TEXT,
      brokerage    TEXT,
      company      TEXT,
      rating_from  TEXT,
      rating_to    TEXT,
      target_from  NUMERIC(10,2),
      target_to    NUMERIC(10,2),
      ticker       TEXT,
      time         TIMESTAMPTZ,
      spot_price   NUMERIC(14,4),  -- nueva columna
      loaded_at    TIMESTAMPTZ NOT NULL
    )`)
	return err
}

// EnsureHistoryTable crea la tabla de historial con spot_price.
func EnsureHistoryTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS StockInfoHistory (
      action       TEXT,
      brokerage    TEXT,
      company      TEXT,
      rating_from  TEXT,
      rating_to    TEXT,
      target_from  NUMERIC(10,2),
      target_to    NUMERIC(10,2),
      ticker       TEXT,
      time         TIMESTAMPTZ,
      spot_price   NUMERIC(14,4),  -- nueva columna
      archived_at  TIMESTAMPTZ DEFAULT now()
    )`)
	return err
}

// LastLoadTime devuelve la última marca de tiempo del registro cargado (loaded_at).
func LastLoadTime(ctx context.Context, conn *pgx.Conn) (time.Time, error) {
	var t *time.Time
	err := conn.QueryRow(ctx, `SELECT max(loaded_at) FROM StockInfo`).Scan(&t)
	if err != nil {
		return time.Time{}, fmt.Errorf("obteniendo última carga: %w", err)
	}
	if t == nil {
		// nunca se ha cargado nada
		return time.Time{}, nil
	}
	return *t, nil
}

// ArchiveCurrent copia todos los datos de StockInfo a StockInfoHistory, incluyendo spot_price.
func ArchiveCurrent(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    INSERT INTO StockInfoHistory (
      action, brokerage, company, rating_from, rating_to,
      target_from, target_to, ticker, time, spot_price
    )
    SELECT
      action, brokerage, company, rating_from, rating_to,
      target_from, target_to, ticker, time, spot_price
    FROM StockInfo
    `)
	return err
}

// TruncateCurrent vacía la tabla principal StockInfo.
func TruncateCurrent(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `TRUNCATE TABLE StockInfo`)
	return err
}

// InsertItems inserta una lista de Item en StockInfo, ahora incluyendo SpotPrice.
func InsertItems(ctx context.Context, conn *pgx.Conn, items []Item, loadedAt time.Time) error {
	for _, it := range items {
		_, err := conn.Exec(ctx, `
        INSERT INTO StockInfo
          (action, brokerage, company, rating_from, rating_to,
           target_from, target_to, ticker, time, spot_price, loaded_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
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
			it.SpotPrice,
			loadedAt,
		)
		if err != nil {
			return fmt.Errorf("al insertar %v: %w", it, err)
		}
	}
	return nil
}
