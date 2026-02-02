package main

import (
	"thinkdrop-backend/internal/bootstrap"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/internal/config/redis"
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/router"
	"thinkdrop-backend/migrations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// → App entry point (starts server)
func main() {

	//->database connection
	db := database.Connection()

	//-> redis conenction
	Redis := redis.NewRedisClient()
	authmiddileware.AuthenticateMiddileware(Redis)

	//-> DB migrations
	migrations.Migrations(db)

	//-> the autentication initing happens here
	Authcontrollers := bootstrap.InitAuth(db, Redis)
	InterestControllers := bootstrap.InitInterest(db)
	PostController := bootstrap.InitPost(db, Redis)
	ProfileController := bootstrap.InitProfile(db)

	//->fibre engine
	app := fiber.New()
	app.Use(logger.New())

	//->pass the engine and controllers for handling the routes
	router.UserRoutes(app, Redis, Authcontrollers, InterestControllers)
	router.OTPRouter(app, Authcontrollers)
	router.PostRoutes(app, PostController, Redis)
	router.ProfileRoute(app, Redis, ProfileController)

	//-> PORT of server
	app.Listen(":8000")
}
