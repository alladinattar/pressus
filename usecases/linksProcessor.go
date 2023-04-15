package usecases

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	"github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func (s *service) ProcessLinks() {
	msgs := make(chan amqp091.Delivery)
	go s.repo.GetTasks(msgs)

	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	for msg := range msgs {
		task := &presenters.ArticleObj{}
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Error("Failed unmarshall task: ", err.Error())
		}
		log.Info("Received from tasks: ", task.Title)
		requestString := fmt.Sprintf("%s%s", s.GetEnv().Config.Parser.DefaultRoute, task.Link)
		var resp []byte
		_, body, err := client.Get(requestString).Get(resp, requestString)
		if err != nil {
			log.Error("Failed request article body: ", err.Error())
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			log.Error("Failed parse article body: ", err.Error())
		}

		doc.Find(".article-body").Each(func(i int, sel *goquery.Selection) {
			s.repo.PushArticleBody(sel.Text())
		})
		msg.Ack(true)
	}
}
