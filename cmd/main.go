package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	internals "github.com/nationpulse-bff/internal/server"
	"github.com/nationpulse-bff/internal/store"
)

//func run(ctx context.Context) {
//ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
//}

func main() {
	ctx := context.Background()
	// Load environment variables
	_ = godotenv.Load()
	for _, k := range []string{"ACCESS_SECRET", "REFRESH_SECRET"} {
		if os.Getenv(k) == "" {
			log.Fatalf("%s not set", k)
		}
	}
	// Create redis client and store in context
	rds := store.NewRedis()
	db := store.NewPgClient()
	ctx = context.WithValue(ctx, "redisClient", rds)
	ctx = context.WithValue(ctx, "dbClient", db)
	// Start a HTTP server
	srv := internals.NewServer(ctx)

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
