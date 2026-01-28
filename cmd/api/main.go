package main

import (
	"thinkdrop-backend/internal/bootstrap"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/internal/config/redis"
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	otprouter "thinkdrop-backend/internal/router/otpRouter"
	userrouter "thinkdrop-backend/internal/router/userRouter"
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
	controllers := bootstrap.InitAuth(db, Redis)
	OtpControllers := bootstrap.InitOTP(Redis)

	//->fibre engine
	app := fiber.New()
	app.Use(logger.New())

	//->pass the engine and controllers for handling the routes
	userrouter.UserRoutes(app, controllers)
	otprouter.OTPRouter(app, OtpControllers)

	//-> PORT of server
	app.Listen(":8000")
}
