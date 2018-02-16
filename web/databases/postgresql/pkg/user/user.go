package user

import (
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/postgres"
	"errors"
)

const (
	tableQuery = `
	CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		email text NOT NULL,
		password varchar(60) NOT NULL,
		permissions int8 NOT NULL
	)`

	getQuery = `
	SELECT * FROM users WHERE email = $1
	`
)

const (
	regular = iota
	moderator = iota
	administrator = iota
)

type User struct {
	ID int `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	Permissions int `json"permissions"`
}

func createTable() (error) {
	db, err := postgres.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(tableQuery); err != nil {
		return err
	}

	return nil
}

// Function NewInstance intakes a pointer to a User struct
// and returns the pointer back as well as an error if
// applicable. The reason for this function is to ensure
// that the users table exists before querying the database
// by use of the unexported createTable function.
func NewInstance(u *User) (*User, error) {
	return u, createTable()
}

// Function Get intakes an email as a string and attempts to
// get the user attached to that email.
func Get(email string) (*User, error) {
	if valid := validateEmail(email); !valid {
		return nil, ErrEmailInvalid
	}

	if err := createTable(); err != nil {
		return nil, err
	}

	db, err := postgres.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var u User

	err = db.QueryRow(getQuery, email).Scan(&u.ID, &u.Email, &u.Password, &u.Permissions)
	if err != nil {
		return nil, err
	}

	return &u, nil
}