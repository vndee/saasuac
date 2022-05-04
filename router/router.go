package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vndee/saasuac/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Healthcheck)

	auth := app.Group("/auth", logger.New())
	auth.Post("/register/", handler.Register)
}
