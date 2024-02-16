package userHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/internal/repository"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/internal/services"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/models"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/jwt"
)

func Login(c *fiber.Ctx) error {
	type LoginInfo struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var user *models.User
	var loginInfo LoginInfo
	var err error

	if err = c.BodyParser(&loginInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Bad Request", "data": err})
	}

	if loginInfo.Email == "" {
		if user, err = repository.GetUserByUsername(loginInfo.Username, loginInfo.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User Not Found"})
		}
	} else {
		if user, err = repository.GetUserByEmail(loginInfo.Email, loginInfo.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "User Not Found"})
		}
	}
	var cookie *fiber.Cookie

	if cookie, err = jwt.CreateToken(user.Username); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Token create failed",
			"data":    err.Error(),
		})
	}
	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": cookie.Value,
		"user": struct {
			Email    string `json:"email"`
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
		}})
}

func Register(c *fiber.Ctx) error {

	user := new(models.User)
	var err error

	if err = c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User register failed",
			"data":    err.Error(),
		})
	}

	err = services.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "User register failed",
			"data":    err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created",
	})
}
