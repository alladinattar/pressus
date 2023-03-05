package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func GetFlows(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(c.Route().Params)
		flows, err := service.GetFlows()
		if err != nil {
			response := presenters.ResponseStruct{
				Data:   nil,
				Status: "failed",
			}
			log.Error("Failed get flows, error: ", err.Error())
			return c.Status(fiber.StatusOK).JSON(response)
		}
		response := presenters.GetFlowsResp{
			Data:   flows,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}

func GetArticlesByFlow(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println(c.Route().Params)
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
