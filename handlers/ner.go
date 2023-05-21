package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func ExtractEntities(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Query("id")
		if id == "" {
			log.Println("Empty id")
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			return c.Status(fiber.StatusBadRequest).JSON(response)
		}
		entities, err := service.ExtractEntities(id)
		if err != nil {
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			log.Error("Failed extract entities from article, error: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(response)
		}
		response := presenters.ResponseStructInterface{
			Data:   entities,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
