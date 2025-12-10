package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PgClient struct{ Client *sql.DB }

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"password"`
}

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

	//sqlStatement := `SELECT * FROM get_perfgrowthpopulation_dashboard(2025, 5)`
	//result, err := db.Query(sqlStatement)
	//if err != nil {
	//	fmt.Printf("Error occured while querying database: %s", err)
	//	panic(err)
	//}
	//defer result.Close()
	//fmt.Println("Query executed successfully")
	//fmt.Printf("%v", result)

	return &PgClient{Client: db}
}

func (pg *PgClient) Close() {
	pg.Client.Close()
}

func (pg *PgClient) GetUser(user *User) (*User, error) {
	sqlStatement := `SELECT * FROM get_user($1, $2);`
	row := pg.Client.QueryRow(sqlStatement, user.Name, user.Email)
	//var name, email string
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	fmt.Printf("Result: id: %s, user:%s, email:%s \n ", user.ID, user.Name, user.Email)

	if err != nil {
		return nil, err
	}

	// user = &User{
	// 	Name: result.Name,
	// }
	return user, nil
}
