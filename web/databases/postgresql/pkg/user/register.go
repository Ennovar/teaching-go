package user

import (
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/postgres"
	"regexp"
	"errors"
)

const registerQuery = `
INSERT INTO users (
	email, password, permissions
) VALUES (
	$1, $2, $3
) RETURNING id`

var (
	ErrEmailInvalid = errors.New("invalid email address")
	ErrPassLength = errors.New("password cannot be less than 8 characters")
	ErrPermInvalid = errors.New("user assigned invalid permission")
)

func validateEmail(email string) bool {
	rx := regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")

	return rx.MatchString(email)
}

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


	return nil
}