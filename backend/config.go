package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Carga .env si existe
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No se pudo cargar .env, usando variables de entorno del sistema")
	}
}

type APIConfig struct {
	URL    string
	Bearer string
}

func LoadAPIConfig() APIConfig {
	url := os.Getenv("URL_APIDATA")
	bearer := os.Getenv("BEARER_APIDATA")
	if url == "" || bearer == "" {
		log.Fatal("Debe definir URL_APIDATA y BEARER_APIDATA en el entorno")
	}
	return APIConfig{URL: url, Bearer: bearer}
}

type DBConfig struct {
	User, Pass, Host, Port, Name, SSLMode string
}

func LoadDBConfig() DBConfig {
	cfg := DBConfig{
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
	return cfg
}
