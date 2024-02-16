package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/database"
	"github.com/hsyntzgl/jwt-authentication-backend-with-middleware.git/pkg/router"
)

func main() {
	app := fiber.New()

	//Database Functions
	database.ConnectDB()
	database.Migrate()

	//Setup routes
	router.SetupRoutes(app)

	app.Listen(":3000")
}
