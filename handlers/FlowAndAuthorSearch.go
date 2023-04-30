package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func SearchByFlowAndAuthor(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flow := c.Query("flow")
		author := c.Query("author")
		articles, err := service.SearchByFlowAndAuthor(flow, author)
		if err != nil {
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			log.Error("Failed search articles, error: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
		response := presenters.ResponseStruct{
			Data:   articles,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
