package store

import (
	"database/sql"
	"fmt"
	"os"
)

type PgClient struct{ pgClient *sql.DB }

func NewPgClient() *PgClient {
	dbName := os.Getenv("PG_DB_NAME")
	dbUser := os.Getenv("PG_DB_USER")
	dbPass := os.Getenv("PG_DB_PASSWORD")

	connStr := fmt.Sprintf("user=%s dbName=%s password=%s sslmode=disable", dbUser, dbName, dbPass)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("Error occured while connecting database: %s", err)
	}
	return &PgClient{pgClient: db}
}
