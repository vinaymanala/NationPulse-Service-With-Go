package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/nationpulse-bff/internal/config"
	"github.com/nationpulse-bff/internal/kafka"
	internals "github.com/nationpulse-bff/internal/server"
	"github.com/nationpulse-bff/internal/store"
	"github.com/nationpulse-bff/internal/utils"
)

//func run(ctx context.Context) {
//ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
//}

func main() {
	ctx := context.Background()
	// Load environment variables from .env for local development
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load; relying on environment variables")
	}

	// Load configuration from environment
	cfg := config.Load()

	// for _, k := range []string{"ACCESS_SECRET", "REFRESH_SECRET"} {
	// 	if os.Getenv(k) == "" {
	// 		log.Fatalf("%s not set", k)
	// 	}
	// }
	// Create redis and postgres store in context
	rds := store.NewRedis(cfg)
	db := store.NewPgClient(ctx, cfg)
	k := kafka.NewKafka(ctx, cfg)

	configs := &utils.Configs{
		Db:      db,
		Cache:   rds,
		Context: ctx,
		Kafka:   k,
		Cfg:     cfg,
	}

	defer rds.Client.Close()
	defer db.Client.Close()

	// Start a HTTP server
	srv := internals.NewServer(configs)

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: srv,
	}
	fmt.Printf("Starting up..")

	log.Printf("listening to %s\n", httpServer.Addr)
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("error listening and serving: %s\n", err)
		os.Exit(1)
	}
}
