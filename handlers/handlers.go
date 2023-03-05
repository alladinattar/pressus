package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func GetFlows(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flows, err := service.GetFlows()
		if err != nil {
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			log.Error("Failed get flows, error: ", err.Error())
			return c.Status(fiber.StatusOK).JSON(response)
		}
		response := presenters.ResponseStruct{
			Data:   flows,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func GetArticlesByFlow(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		flow := c.Params("flow_name")
		articles, err := service.GetArticlesByFlow(flow)
		if err != nil {
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			log.Error("Failed get flows, error: ", err.Error())
			return c.Status(fiber.StatusOK).JSON(response)
		}
		response := presenters.ResponseStruct{
			Data:   articles,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
