package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/pressus/usecases"
	log "github.com/sirupsen/logrus"
)

func GetData(service usecases.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, _ := service.GetEnv().SessionStore.Get(c)
		log.Info("Session fresh 2: ", sess.Fresh())
		domains := service.GetData()
		response := presenters.ResponseStruct{
			Data:   domains,
			Status: "success",
		}
		return c.Status(fiber.StatusOK).JSON(response)
	}
}
