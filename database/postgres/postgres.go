package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func App() *sql.DB {

	host := os.Getenv("REST_API_HOST")

	dbname := os.Getenv("REST_API_DB")

	user := os.Getenv("REST_API_USER")

	password := os.Getenv("REST_API_PASSWORD")

	port := os.Getenv("REST_API_PG_PORT")

	source := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)

	conn, err := sql.Open("postgres", source)

	if err != nil {
		log.Fatalf(`Database error: %s`, err.Error())
		return nil
	}

	return conn

}
