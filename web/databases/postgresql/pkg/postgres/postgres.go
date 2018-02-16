package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const(
	user = "postgres"
	password = "postgres"
	database = "teaching-go"
)

func Open() (*sql.DB, error) {
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, database)
	return sql.Open("postgres", conn)
}