package routes

import (
	"github.com/gofiber/fiber/v2"
	helloHandlers "github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/internal/handlers/hello"
)

func SetupHelloRoutes(router fiber.Router) {
	hello := router.Group("/hello")
	hello.Get("/", helloHandlers.SayHello)
}
