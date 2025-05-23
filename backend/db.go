package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

// Ensure main table with loaded_at
func EnsureTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS StockInfo (
      action      TEXT,
      brokerage   TEXT,
      company     TEXT,
      rating_from TEXT,
      rating_to   TEXT,
      target_from TEXT,
      target_to   TEXT,
      ticker      TEXT,
      time        TIMESTAMPTZ,
      loaded_at   TIMESTAMPTZ NOT NULL
    )`)
	return err
}

// Ensure history/archive table
func EnsureHistoryTable(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    CREATE TABLE IF NOT EXISTS StockInfoHistory (
      action      TEXT,
      brokerage   TEXT,
      company     TEXT,
      rating_from TEXT,
      rating_to   TEXT,
      target_from TEXT,
      target_to   TEXT,
      ticker      TEXT,
      time        TIMESTAMPTZ,
      archived_at TIMESTAMPTZ DEFAULT now()
    )`)
	return err
}

// LastLoadTime devuelve la última marca de tiempo del registro archivado, manejando NULL
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

// ArchiveCurrent copia los datos de StockInfo en StockInfoHistory
func ArchiveCurrent(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `
    INSERT INTO StockInfoHistory (action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time)
    SELECT action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time FROM StockInfo`)
	return err
}

// TruncateCurrent vacía la tabla principal
func TruncateCurrent(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, `TRUNCATE TABLE StockInfo`)
	return err
}

// InsertItems inserta nuevos ítems con timestamp de carga
func InsertItems(ctx context.Context, conn *pgx.Conn, items []Item, loadedAt time.Time) error {
	for _, it := range items {
		_, err := conn.Exec(ctx, `
        INSERT INTO StockInfo
          (action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time, loaded_at)
        VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
        `, it.Action, it.Brokerage, it.Company, it.RatingFrom, it.RatingTo, it.TargetFrom, it.TargetTo, it.Ticker, it.Time, loadedAt)
		if err != nil {
			return fmt.Errorf("al insertar %v: %w", it, err)
		}
	}
	return nil
}
