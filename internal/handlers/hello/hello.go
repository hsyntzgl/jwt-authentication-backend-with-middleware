package helloHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/jwt"
)

func SayHello(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")
	err := jwt.CheckUserToken(tokenString)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Hello !",
	})
}
