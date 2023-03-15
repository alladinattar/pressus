package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func SaveArticlesByFlow(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flow := c.Params("flow_name")
		result, err := service.SaveArticlesByFlow(flow)

		if err != nil {
			response := presenters.StatusResponseStruct{
				Status: "failed",
			}
			log.Error("Failed get flows, error: ", err.Error())
			return c.Status(fiber.StatusOK).JSON(response)
		}
		if result != true {
			response := presenters.StatusResponseStruct{
				Status: "failed",
			}
			log.Error("Failed save flow, error: ", err.Error())
			return c.Status(fiber.StatusOK).JSON(response)
		}

		response := presenters.StatusResponseStruct{
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)

	}
}
