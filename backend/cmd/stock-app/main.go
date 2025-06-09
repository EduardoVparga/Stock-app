// file name: backend/cmd/stock-analyzer/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"stock-app/internal/apidata"
	"stock-app/internal/config"

	// "stock-app/internal/processor"
	// "stock-app/internal/spotclient"
	"stock-app/internal/storage"
	"stock-app/internal/transport/http"
)

func main() {
	// Crear un contexto principal que se cancela con señales del sistema
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Manejar señales de interrupción para un apagado ordenado
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		log.Println("🛑 Apagado solicitado, finalizando procesos...")
		cancel()
	}()

	// 1. Cargar configuración
	config.LoadEnv("../.env")
	apiCfg := config.LoadAPIConfig()
	dbCfg := config.LoadDBConfig()

	// 2. Inicializar dependencias (capa de almacenamiento y clientes)
	repo, err := storage.NewRepository(ctx, dbCfg.DSN)
	if err != nil {
		log.Fatalf("Error inicializando el repositorio: %v", err)
	}
	defer repo.Close()

	apiDataClient := apidata.NewClient(apiCfg)
	// spotAPIClient := spotclient.NewClient("http://localhost:8000") // URL hardcodeada, podría ir a config

	// 3. Lógica de arranque: actualizar datos si es necesario
	if err := runInitialDataLoad(ctx, repo, apiDataClient); err != nil {
		log.Fatalf("Error durante la carga inicial de datos: %v", err)
	}

	// 4. Iniciar servicios en segundo plano (goroutines)
	// priceScheduler := processor.NewScheduler(repo, spotAPIClient)
	// go priceScheduler.Start(ctx)

	// 5. Iniciar el servidor HTTP (bloqueante)
	server := http.NewServer(repo)
	server.Start(ctx, ":8080")

	log.Println("Aplicación finalizada.")
}

// runInitialDataLoad comprueba si es necesario cargar nuevos datos de la API.
func runInitialDataLoad(ctx context.Context, repo *storage.Repository, apiClient *apidata.Client) error {
	lastLoad, err := repo.GetLastLoadTime(ctx)
	if err != nil {
		// Asumimos que la tabla podría estar vacía o la función no existir aún, no es fatal
		log.Printf("No se pudo obtener la última carga (puede ser la primera vez): %v", err)
	}

	if !lastLoad.IsZero() && time.Since(lastLoad) < time.Hour {
		log.Println("Última carga hace menos de una hora. Saltando actualización.")
		return nil
	}

	log.Println("Iniciando actualización de datos desde la API...")

	if err := repo.ExecuteProcedure(ctx, "ArchiveAndTruncate"); err != nil {
		return fmt.Errorf("error en ArchiveAndTruncate: %w", err)
	}

	items, err := apiClient.FetchAll(ctx)
	if err != nil {
		return fmt.Errorf("error al extraer datos: %w", err)
	}
	log.Printf("Registros obtenidos: %d\n", len(items))

	if err := repo.InsertStockInfo(ctx, items, time.Now()); err != nil {
		return fmt.Errorf("error insertando datos: %w", err)
	}

	if err := repo.ExecuteProcedure(ctx, "sync_symbols"); err != nil {
		return fmt.Errorf("error en sync_symbols: %w", err)
	}

	log.Println("✅ Datos actualizados y almacenados en StockInfo")
	return nil
}
