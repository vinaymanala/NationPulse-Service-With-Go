package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PgClient struct{ Client *pgxpool.Pool }

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewPgClient(ctx context.Context) *PgClient {
	connStr := "postgres://postgres:postgres@localhost:5432/nationPulseDB?sslmode=disable"
	fmt.Println(connStr)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Printf("Error occured while connecting database: %s\n", err)
		panic(err)
	}
	fmt.Println("Connected to Postgres database successfully")
	return &PgClient{Client: pool}
}

func (pg *PgClient) GetUser(ctx context.Context, user *User) (*User, error) {
	sqlStatement := `SELECT * FROM get_user($1, $2);`
	row := pg.Client.QueryRow(ctx, sqlStatement, user.Name, user.Email)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	fmt.Printf("Result: id: %s, user:%s, email:%s \n ", user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}
