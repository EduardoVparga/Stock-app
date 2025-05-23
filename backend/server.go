package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jackc/pgx/v4"
	"github.com/rs/cors"
)

// ServidorHTTP expone un endpoint GET /stocks con la data de la tabla StockInfo
func ServidorHTTP(ctx context.Context, conn *pgx.Conn) {
	mux := http.NewServeMux()
	mux.HandleFunc("/stocks", func(w http.ResponseWriter, r *http.Request) {
		rows, err := conn.Query(ctx, `
            SELECT action, brokerage, company, rating_from, rating_to,
                   target_from, target_to, ticker, time
            FROM StockInfo ORDER BY time DESC
        `)
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			log.Println("DB error:", err)
			return
		}
		defer rows.Close()

		var stocks []Item
		for rows.Next() {
			var it Item
			err := rows.Scan(&it.Action, &it.Brokerage, &it.Company, &it.RatingFrom, &it.RatingTo,
				&it.TargetFrom, &it.TargetTo, &it.Ticker, &it.Time)
			if err != nil {
				http.Error(w, "Error al leer los datos", http.StatusInternalServerError)
				log.Println("Scan error:", err)
				return
			}
			stocks = append(stocks, it)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stocks)
	})

	// Configura el middleware CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Reemplaza con el origen de tu app Vue
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	log.Println("🌐 API disponible en http://localhost:8080/stocks")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
