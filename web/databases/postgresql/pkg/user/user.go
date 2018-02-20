package user

import (
	"errors"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/postgres"
)

const (
	tableQuery = `CREATE TABLE IF NOT EXISTS users (
		id serial PRIMARY KEY,
		email text NOT NULL,
		password varchar(60) NOT NULL,
		permissions int8 NOT NULL
	)`

	getQueryEmail = "SELECT * FROM users WHERE email = $1 LIMIT 1"
	getQueryID    = "SELECT * FROM users WHERE ID = $1 LIMIT 1"
)

const (
	regular = iota
	moderator
	administrator
)

var (
	ErrInvalidIdentifier = errors.New("the identifier used is not valid")
)

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Permissions int    `json:"permissions"`
	SessionID   string `json:"sessionID,omitempty"`
}

func createTable() error {
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

// Function Get intakes either an email of type string or
// an ID of type int and gets the user based off of the
// identifier given.
func Get(identifier interface{}) (*User, error) {
	if err := createTable(); err != nil {
		return nil, err
	}

	db, err := postgres.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var u User

	email, ok := identifier.(string)
	if !ok {
		id, ok := identifier.(int)
		if !ok {
			return nil, ErrInvalidIdentifier
		}

		err := db.QueryRow(getQueryID, id).Scan(&u.ID, &u.Email, &u.Password, &u.Permissions, &u.SessionID)
		if err != nil {
			return nil, err
		}

		return &u, nil
	}

	if valid := validateEmail(email); !valid {
		return nil, ErrEmailInvalid
	}

	err = db.QueryRow(getQueryEmail, email).Scan(&u.ID, &u.Email, &u.Password, &u.Permissions, &u.SessionID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) GetStatus() string {
	switch u.Permissions {
	case regular:
		return "regular"
	case moderator:
		return "moderator"
	case administrator:
		return "administrator"
	default:
		return ""
	}
}
