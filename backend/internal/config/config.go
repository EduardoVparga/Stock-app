package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	if err := godotenv.Load(path); err != nil {
		log.Println("No se pudo cargar el archivo .env, usando variables de entorno del sistema")
	}
}

type API struct {
	URL    string
	Bearer string
}

func LoadAPIConfig() API {
	url := os.Getenv("URL_APIDATA")
	bearer := os.Getenv("BEARER_APIDATA")
	if url == "" || bearer == "" {
		log.Fatal("Debe definir URL_APIDATA y BEARER_APIDATA en el entorno")
	}
	return API{URL: url, Bearer: bearer}
}

type DB struct {
	DSN string
}

func LoadDBConfig() DB {
	cfg := struct {
		User, Pass, Host, Port, Name, SSLMode string
	}{
		User:    os.Getenv("DB_USER"),
		Pass:    os.Getenv("DB_PASS"),
		Host:    os.Getenv("DB_HOST"),
		Port:    os.Getenv("DB_PORT"),
		Name:    os.Getenv("DB_NAME"),
		SSLMode: os.Getenv("DB_SSLMODE"),
	}
	if cfg.User == "" || cfg.Pass == "" || cfg.Host == "" || cfg.Port == "" || cfg.Name == "" {
		log.Fatal("Faltan variables de entorno para la BD (DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)")
	}
	if cfg.SSLMode == "" {
		cfg.SSLMode = "require"
	}
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode,
	)
	return DB{DSN: dsn}
}
