package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/pressus/config"
	"github.com/pressus/repository/queue"
	"github.com/pressus/repository/search-engine"
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
		AppName: "pressus",
	})

	app.Use(pprof.New())
	app.Use(logger.New())

	queueConnection, channel, err := QueueConnection(env)
	defer queueConnection.Close()
	defer channel.Close()
	if err != nil {
		log.Fatalf("Failed connect to queue: %s", err.Error())
	}

	repo := queue.NewQueueRepo(queueConnection, channel)
	searchEngine := search.NewEngineRepo(queueConnection, channel)
	service := usecases.NewService(env, repo, searchEngine)

	go service.ProcessLinks()

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
