package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type PgClient struct{ Client *pgx.Conn }

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"password"`
}

func NewPgClient(ctx context.Context) *PgClient {
	//dbHost := os.Getenv("PG_DB_HOST")
	//dbName := os.Getenv("PG_DB_NAME")
	//dbUser := os.Getenv("PG_DB_USER")
	//dbPass := os.Getenv("PG_DB_PASSWORD")

	//connStr := fmt.Sprintf("host=localhost post=5433 user=%s dbName=%s password=%s sslmode=disable", dbUser, dbName, dbPass)
	connStr := "postgres://postgres:postgres@localhost:5432/nationPulseDB?sslmode=disable"
	//connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":5432/" + dbName + "?sslmode=disable"
	fmt.Println(connStr)
	conn, err := pgx.Connect(ctx, connStr)
	//fmt.Println(conn)
	if err != nil {
		defer conn.Close(ctx)
		fmt.Printf("Error occured while connecting database: %s", err)
		panic(err)
	}
	fmt.Println("Connected to Postgres database successfully")
	return &PgClient{Client: conn}
}

func (pg *PgClient) GetUser(ctx context.Context, user *User) (*User, error) {
	sqlStatement := `SELECT * FROM get_user($1, $2);`
	row := pg.Client.QueryRow(ctx, sqlStatement, user.Name, user.Email)
	//var name, email string
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	fmt.Printf("Result: id: %s, user:%s, email:%s \n ", user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
