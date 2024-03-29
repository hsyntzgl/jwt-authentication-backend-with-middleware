package jwt

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my_secret_key")

func CreateToken(username string) (*fiber.Cookie, error) {

	expireAt := time.Now().Add(time.Minute * 5)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expireAt.Unix(),
	})

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  expireAt,
		HTTPOnly: true,
	}
	return &cookie, nil
}
func CheckUserToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		err := errorHandler("Token is invalid")
		return err
	}
	return nil
}

func errorHandler(message string) error {
	err := errors.New(message)
	return err
}
