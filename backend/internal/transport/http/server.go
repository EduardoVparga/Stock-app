// file name: backend/internal/transport/http/server.go
package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"stock-app/internal/storage"

	"github.com/rs/cors"
)

// Server expone la API a través de HTTP.
type Server struct {
	repo *storage.Repository
}

// NewServer crea una instancia del servidor HTTP.
func NewServer(repo *storage.Repository) *Server {
	return &Server{repo: repo}
}

// Start inicia el servidor y escucha en el puerto especificado.
func (s *Server) Start(ctx context.Context, addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/stocks", s.handleGetStocks(ctx))

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)

	log.Printf("🌐 API disponible en http://localhost%s/stocks", addr)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("Error al iniciar el servidor HTTP: %v", err)
	}
}

func (s *Server) handleGetStocks(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stocks, err := s.repo.GetRankedStocks(r.Context()) // Usar el contexto de la request
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			log.Println("DB error:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(stocks); err != nil {
			log.Printf("Error al codificar respuesta JSON: %v", err)
		}
	}
}
