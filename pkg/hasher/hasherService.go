package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

const DefaultCost = 10

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)

	return string(bytes), err
}
func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
