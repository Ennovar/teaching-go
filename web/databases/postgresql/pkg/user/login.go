package user

import (
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/encryption"
	"github.com/Ennovar/teaching-go/web/databases/postgresql/pkg/postgres"
)

const loginQuery = `UPDATE users
SET session_id=$1
WHERE ID=$2`

// Function login takes a user struct and will log the
// user into the system by generating a session ID to
// be stored in the users cookie. The session struct
// will be returned upon successful login. To get a user
// struct to utilize the login function, use the user.Get()
// function.
func (u *User) Login(password string) (string, error) {
	if err := encryption.CheckPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", err
	}

	db, err := postgres.Open()
	if err != nil {
		return "", err
	}
	defer db.Close()

	sessionID := encryption.RandomString(16)

	_, err = db.Exec(loginQuery, sessionID, u.ID)
	if err != nil {
		return "", err
	}

	return sessionID, nil
}
