package router

import (
	"thinkdrop-backend/internal/modules/admin/delivery"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)


func AdminRoutes(app *fiber.App, rds *redis.Client, AdminModule *delivery.AdminController){
	
}