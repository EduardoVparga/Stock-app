package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

// APIConfig almacena la configuración de la API.
type APIConfig struct {
	URL    string
	Bearer string
}

// Item representa un registro individual de la API.
type Item struct {
	Action     string    `json:"action"`
	Brokerage  string    `json:"brokerage"`
	Company    string    `json:"company"`
	RatingFrom string    `json:"rating_from"`
	RatingTo   string    `json:"rating_to"`
	TargetFrom string    `json:"target_from"`
	TargetTo   string    `json:"target_to"`
	Ticker     string    `json:"ticker"`
	Time       time.Time `json:"time"`
}

// APIResponse mapea la respuesta JSON paginada.
type APIResponse struct {
	Items    []Item `json:"items"`
	NextPage string `json:"next_page"`
}

// loadEnv carga el archivo .env (si existe) y luego confía en las variables de entorno.
func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar .env, usando variables de entorno del sistema")
	}
}

// loadAPIConfig lee URL_APIDATA y BEARER_APIDATA de las variables de entorno.
func loadAPIConfig() APIConfig {
	url := os.Getenv("URL_APIDATA")
	bearer := os.Getenv("BEARER_APIDATA")
	if url == "" || bearer == "" {
		log.Fatal("Debe definir URL_APIDATA y BEARER_APIDATA en el entorno")
	}
	return APIConfig{URL: url, Bearer: bearer}
}

// loadDBConfig lee la configuración de la base de datos de las variables de entorno.
func loadDBConfig() (user, pass, host, port, name, sslmode string) {
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASS")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	name = os.Getenv("DB_NAME")
	sslmode = os.Getenv("DB_SSLMODE")
	if user == "" || pass == "" || host == "" || port == "" || name == "" {
		log.Fatal("Faltan variables de entorno para la BD (DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)")
	}
	if sslmode == "" {
		sslmode = "require"
	}
	return
}

// fetchAllItems recorre todas las páginas de la API y retorna todos los items.
func fetchAllItems(cfg APIConfig) ([]Item, error) {
	var all []Item
	next := ""
	for {
		req, err := http.NewRequest("GET", cfg.URL, nil)
		if err != nil {
			return nil, fmt.Errorf("error creando request: %w", err)
		}
		if next != "" {
			q := req.URL.Query()
			q.Add("next_page", next)
			req.URL.RawQuery = q.Encode()
		}
		req.Header.Set("Authorization", "Bearer "+cfg.Bearer)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error en petición HTTP: %w", err)
		}
		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("error leyendo respuesta: %w", err)
		}
		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status %s: %s", resp.Status, string(body))
		}

		var r APIResponse
		if err := json.Unmarshal(body, &r); err != nil {
			return nil, fmt.Errorf("error parseando JSON: %w", err)
		}

		all = append(all, r.Items...)
		if r.NextPage == "" {
			break
		}
		next = r.NextPage
	}
	return all, nil
}

// ensureTable crea la tabla StockInfo si no existe.
func ensureTable(ctx context.Context, conn *pgx.Conn) error {
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
	  time        TIMESTAMPTZ
	)`)
	return err
}

// insertItems inserta los registros en la tabla StockInfo.
func insertItems(ctx context.Context, conn *pgx.Conn, items []Item) error {
	for _, it := range items {
		_, err := conn.Exec(ctx, `
		INSERT INTO StockInfo
		  (action, brokerage, company, rating_from, rating_to, target_from, target_to, ticker, time)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
		`, it.Action, it.Brokerage, it.Company, it.RatingFrom, it.RatingTo, it.TargetFrom, it.TargetTo, it.Ticker, it.Time)
		if err != nil {
			return fmt.Errorf("al insertar registro %v: %w", it, err)
		}
	}
	return nil
}

func main() {
	// 1) Cargar variables de entorno
	loadEnv()

	// 2) Leer configuración API y BD
	apiCfg := loadAPIConfig()
	dbUser, dbPass, dbHost, dbPort, dbName, dbSSL := loadDBConfig()

	// 3) Construir DSN y conectar a CockroachDB
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, dbSSL,
	)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("No se pudo conectar a la BD: %v", err)
	}
	defer conn.Close(ctx)

	// 4) Crear tabla si no existe
	if err := ensureTable(ctx, conn); err != nil {
		log.Fatalf("Error creando tabla StockInfo: %v", err)
	}

	// 5) Extraer todos los items de la API
	items, err := fetchAllItems(apiCfg)
	if err != nil {
		log.Fatalf("Error al extraer datos de la API: %v", err)
	}
	fmt.Printf("Registros obtenidos: %d\n", len(items))

	// 6) Insertar en la tabla StockInfo
	if err := insertItems(ctx, conn, items); err != nil {
		log.Fatalf("Error insertando datos en la BD: %v", err)
	}

	fmt.Println("✅ Todos los datos se han almacenado en la tabla StockInfo")
}
