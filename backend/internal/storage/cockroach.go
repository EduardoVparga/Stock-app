// file name: backend/internal/storage/cockroach.go
package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"stock-app/internal/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, dsn string) (*Repository, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar a la BD: %w", err)
	}
	return &Repository{pool: pool}, nil
}

func (r *Repository) Close() {
	r.pool.Close()
}

func (r *Repository) InsertStockInfo(ctx context.Context, items []domain.StockInfo, loadedAt time.Time) error {
	for _, it := range items {
		_, err := r.pool.Exec(ctx, `
        INSERT INTO StockInfo
          (action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time, loaded_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
        `,
			it.Action, it.Brokerage, it.Company, it.RatingFrom, it.RatingTo,
			float64(it.TargetFrom), float64(it.TargetTo), it.Ticker, it.Time, loadedAt,
		)
		if err != nil {
			return fmt.Errorf("al insertar %v: %w", it, err)
		}
	}
	return nil
}

func (r *Repository) UpsertSpotPrice(ctx context.Context, ticker string, price float64) error {
	_, err := r.pool.Exec(ctx, `
        UPSERT INTO SpotPrices (ticker, price, loaded_at)
        VALUES ($1, $2, now())
    `, ticker, price)
	return err
}

func (r *Repository) GetLastLoadTime(ctx context.Context) (time.Time, error) {
	var lastLoad time.Time
	err := r.pool.QueryRow(ctx, `SELECT LastLoadTime()`).Scan(&lastLoad)
	return lastLoad, err
}

func (r *Repository) GetActiveTickers(ctx context.Context) ([]string, error) {
	rows, err := r.pool.Query(ctx, "SELECT ticker FROM get_active_tickers()")
	if err != nil {
		return nil, err
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
	return tickers, rows.Err()
}

func (r *Repository) GetRankedStocks(ctx context.Context) ([]map[string]interface{}, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT to_jsonb(stock_scores) AS data
         FROM stock_scores
         ORDER BY rank;`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var raw []byte
		if err := rows.Scan(&raw); err != nil {
			return nil, err
		}

		var m map[string]interface{}
		if err := json.Unmarshal(raw, &m); err != nil {
			return nil, err
		}
		results = append(results, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (r *Repository) ExecuteProcedure(ctx context.Context, procedureName string) error {

	validProcedures := map[string]bool{
		"ArchiveAndTruncate": true,
		"sync_symbols":       true,
	}
	if !validProcedures[procedureName] {
		return fmt.Errorf("procedimiento no válido: %s", procedureName)
	}

	_, err := r.pool.Exec(ctx, fmt.Sprintf(`CALL %s()`, procedureName))
	return err
}
