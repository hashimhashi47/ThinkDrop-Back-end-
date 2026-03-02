package main

import (
	"os"
	"thinkdrop-backend/internal/bootstrap"
	"thinkdrop-backend/internal/config/database"
	"thinkdrop-backend/internal/config/redis"
	authmiddileware "thinkdrop-backend/internal/middleware/authMiddileware"
	"thinkdrop-backend/internal/router"
	"thinkdrop-backend/migrations"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// → App entry point (starts server)
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	//->database connection
	db := database.Connection()

	//-> redis conenction
	Redis := redis.NewRedisClient()
	authmiddileware.AuthenticateMiddileware(Redis)

	//-> DB migrations
	migrations.Migrations(db)

	//-> the autentication initing happens here
	AdminControllers, Adminservice := bootstrap.InitAdmin(db)
	Authcontrollers := bootstrap.InitAuth(db, Redis, Adminservice)
	InterestControllers := bootstrap.InitInterest(db)
	PostController := bootstrap.InitPost(db, Redis, Adminservice)
	ProfileController := bootstrap.InitProfile(db, Adminservice)
	RewardControllers := bootstrap.InitRewards(db, Adminservice)
	ChatControllers := bootstrap.InitChat(db)

	//->fibre engine
	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://thinkdrop-nu.vercel.app",
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
	router.ChatRoutes(app, Redis, ChatControllers)
	router.AdminRoutes(app, Redis, AdminControllers)

	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}

}
