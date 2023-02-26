package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/handlers"
	"github.com/pressus/usecases"
)

func SetupRoutes(app fiber.Router, service usecases.Service) {
	api := app.Group("api/v1")

	api.Get("data", handlers.GetData(service))
}
