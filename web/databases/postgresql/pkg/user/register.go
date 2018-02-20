package user

import (
	"errors"
	"regexp"

	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/encryption"
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/postgres"
	"database/sql"
)

const registerQuery = `
INSERT INTO users (
	email, password, permissions
) VALUES (
	$1, $2, $3
) RETURNING id`

var (
	ErrEmailInvalid = errors.New("invalid email address")
	ErrPassLength   = errors.New("password cannot be less than 8 characters")
	ErrPermInvalid  = errors.New("user assigned invalid permission")
	ErrEmailExists  = errors.New("a user already exists with the given email")
)

func validateEmail(email string) bool {
	rx := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")

	return rx.MatchString(email)
}

// Function Register is attached to the user struct and
// will attempt to register a new user. To get the user
// struct for registering, utilize the function
// user.NewInstance().
func (u *User) Register() error {
	db, err := postgres.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	// Validate Email Address
	if valid := validateEmail(u.Email); !valid {
		return ErrEmailInvalid
	}

	// Validate Password
	if len(u.Password) < 8 {
		return ErrPassLength
	}

	// Validate Permissions
	if u.Permissions < regular || u.Permissions > administrator {
		return ErrPermInvalid
	}

	// Check if email already exists
	_, err = Get(u.Email)
	if err != sql.ErrNoRows {
		return ErrEmailExists
	}

	p, err := encryption.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = db.Exec(registerQuery, u.Email, p, u.Permissions)
	if err != nil {
		return err
	}

	return nil
}
