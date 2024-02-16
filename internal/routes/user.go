package routes

import (
	"github.com/gofiber/fiber/v2"
	userHandlers "github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/internal/handlers/user"
)

func SetupUserRoutes(router fiber.Router) {

	// login := router.Group("/login")
	// register := router.Group("/register")

	router.Post("/login", userHandlers.Login)
	router.Post("/register", userHandlers.Register)
}
