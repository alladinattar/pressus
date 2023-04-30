package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/handlers"
	"github.com/pressus/usecases"
)

func SetupRoutes(app fiber.Router, service usecases.Service) {
	api := app.Group("/api/v1")

	search := api.Group("/search")
	search.Get("/author-flow", handlers.SearchByFlowAndAuthor(service))
	search.Get("/date-flow", handlers.SearchByFlowAndDate(service))

	flows := api.Group("/flows")
	flows.Get("/", handlers.GetFlows(service))
	flows.Get("/:flow_name", handlers.GetArticlesByFlow(service))
	flows.Post("/:flow_name", handlers.SaveArticlesByFlow(service))

}
