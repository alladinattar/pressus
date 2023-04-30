package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func SearchByFlowAndDate(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flow := c.Query("flow")
		from := c.Query("from")
		until := c.Query("until")
		articles, err := service.SearchByFlowAndDate(flow, from, until)
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
