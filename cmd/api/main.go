package main

import (
	"thinkdrop-backend/internal/bootstrap"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/internal/config/redis"
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/modules/chat/websocket"
	"thinkdrop-backend/internal/router"
	"thinkdrop-backend/migrations"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// → App entry point (starts server)
func main() {

	//->database connection
	db := database.Connection()

	//-> redis conenction
	Redis := redis.NewRedisClient()
	authmiddileware.AuthenticateMiddileware(Redis)
	hub := websocket.NewHub()
	go hub.Run()

	//-> DB migrations
	migrations.Migrations(db)

	//-> the autentication initing happens here
	Authcontrollers := bootstrap.InitAuth(db, Redis)
	InterestControllers := bootstrap.InitInterest(db)
	PostController := bootstrap.InitPost(db, Redis)
	ProfileController := bootstrap.InitProfile(db)
	RewardControllers := bootstrap.InitRewards(db)

	//->fibre engine
	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Authorization",
		AllowCredentials: true,
		MaxAge:           int((12 * time.Hour).Seconds()),
	}))

	//->pass the engine and controllers for handling the routes
	router.UserRoutes(app, Redis, Authcontrollers, InterestControllers)
	router.OTPRouter(app, Authcontrollers)
	router.PostRoutes(app, PostController, Redis)
	router.ProfileRoute(app, Redis, ProfileController)
	router.RewardRoutes(app, Redis, db, RewardControllers)

	//-> PORT of server
	app.Listen(":8000")
}
