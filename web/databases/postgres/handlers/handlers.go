package handlers

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Database *sql.DB

func init() {
	var dbName, dbUser, dbPass string
	flag.StringVar(&dbName, "db", "genericForum", "The name of the postgres database being utilized for this application.")
	flag.StringVar(&dbUser, "user", "root", "The name of the postgres user that has sufficient permissions for the database.")
	flag.StringVar(&dbPass, "pass", "root", "The password for the given postgres user.")

	var err error
	conn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPass, dbName)

	Database, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}

	tableTx, err := Database.Begin()
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}

	for _, createQuery := range schema {
		_, err := tableTx.Exec(createQuery)
		if err != nil {
			tableTx.Rollback()
			log.Fatalf("A fatal error has occurred: %v\n", err.Error())
		}
	}

	err = tableTx.Commit()
	if err != nil {
		log.Fatalf("A fatal error has occurred: %v\n", err.Error())
	}
}
