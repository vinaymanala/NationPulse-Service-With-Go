package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PgClient struct{ Client *sql.DB }

func NewPgClient() *PgClient {
	//dbHost := os.Getenv("PG_DB_HOST")
	//dbName := os.Getenv("PG_DB_NAME")
	//dbUser := os.Getenv("PG_DB_USER")
	//dbPass := os.Getenv("PG_DB_PASSWORD")

	//connStr := fmt.Sprintf("host=localhost post=5433 user=%s dbName=%s password=%s sslmode=disable", dbUser, dbName, dbPass)
	connStr := "postgres://postgres:postgres@localhost:5432/nationPulseDB?sslmode=disable"
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		defer db.Close()
		fmt.Printf("Error occured while connecting database: %s", err)
		panic(err)
	}
	fmt.Println("Connected to Postgres database successfully")

	sqlStatement := `SELECT * FROM get_perfgrowthpopulation_dashboard(2025, 5)`
	result, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Printf("Error occured while querying database: %s", err)
		panic(err)
	}
	//defer result.Close()
	fmt.Println("Query executed successfully\n")
	fmt.Printf("%v", result)

	return &PgClient{Client: db}
}
