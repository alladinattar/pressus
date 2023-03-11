package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/pressus/config"
	"github.com/pressus/routes"
	"github.com/pressus/usecases"
	"log"
)

func Run() {
	env := config.NewEnv()
	err := readConfig(env.Config)
	if err != nil {
		log.Fatalf("Failed read config file %s - ", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "pressus-producer",
	})

	app.Use(pprof.New())
	app.Use(logger.New())

	service := usecases.NewService(env)
	routes.SetupRoutes(app, service)
	log.Fatal(app.Listen(env.Config.Api.Port))
}

func readConfig(cfg *config.Config) error {
	err := config.ReadConfFile(cfg)
	if err != nil {
		return err
	}
	err = config.ReadEnv(cfg)
	if err != nil {
		return err
	}
	return nil
}
