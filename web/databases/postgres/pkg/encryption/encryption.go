package encryption

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567891234567890!@#$%^&*()"

// RandomString function takes an integer length value in and returns a
// random string of that size built from the charset constant.
func RandomString(length int) string {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seed.Intn(len(charset))]
	}

	return string(b)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPassword(hash, plainText []byte) error {
	return bcrypt.CompareHashAndPassword(hash, plainText)
}
