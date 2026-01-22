package main

import (
	"thinkdrop-backend/internal/bootstrap"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/internal/router"
	"thinkdrop-backend/migrations"

	"github.com/gofiber/fiber/v2"
)

// → App entry point (starts server)
func main() {

	//->database connection
	db := database.Connection()

	//-> DB migrations
	migrations.Migrations(db)

	//-> the autentication initing happens here
	controllers := bootstrap.InitAuth(db)

	//->fibre engine
	app := fiber.New()

	//->pass the engine and controllers for handling the routes
	router.UserRoutes(app, controllers)

	//-> PORT of server
	app.Listen(":8000")
}
