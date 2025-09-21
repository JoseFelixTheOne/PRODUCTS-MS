package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JoseFelixTheOne/products-ms/internal/config"
	"github.com/JoseFelixTheOne/products-ms/internal/db"
	httpTransport "github.com/JoseFelixTheOne/products-ms/internal/transport/http"
)

func main() {
	cfg := config.Load()
	gormDB := db.MustConnect(cfg)

	r := httpTransport.NewRouter(cfg, gormDB)

	// Graceful shutdown
	go func() {
		if err := r.Run(":" + cfg.AppPort); err != nil {
			log.Fatalf("server error: %v", err)
		}
	}()

	log.Printf("server running on :%s", cfg.AppPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")
}
