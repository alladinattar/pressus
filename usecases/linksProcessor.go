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
	"time"
)

func (s *service) ProcessLinks() {
	msgs := make(chan amqp091.Delivery)
	go s.repoTasks.GetTasks(msgs)
	client := fiber.Client{
		UserAgent: s.env.Config.Parser.UserAgent,
	}

	for msg := range msgs {
		task := &presenters.ArticleObj{}
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Error("Failed unmarshall task: ", err.Error())
		}
		log.Info("Received task: ", task.Title)
		requestString := fmt.Sprintf("%s%s", s.GetEnv().Config.Parser.DefaultRoute, task.Link)
		var resp []byte
		_, body, err := client.Get(requestString).Timeout(time.Second*5).Get(resp, requestString)
		if err != nil {
			log.Error("Failed request article body: ", err.Error())
			continue
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			log.Error("Failed parse article body: ", err.Error())
			continue
		}

		doc.Find(".article-body").First().Each(func(i int, sel *goquery.Selection) {
			task.Body = sel.Text()
		})
		s.repoTasks.PushArticleToResults(task)
		msg.Ack(true)
	}

}

func (s *service) ProcessLinksFromResultQueue() {
	articles := make(chan amqp091.Delivery)
	go s.repoResult.GetResults(articles)

	for msg := range articles {
		article := &presenters.ArticleObj{}
		err := json.Unmarshal(msg.Body, &article)
		if err != nil {
			log.Error("Failed unmarshall article: ", err.Error())
		}
		log.Info("Received result task: ", article.Title)

		if exist, err := s.searchEngine.IsArticleExist(article.ID); exist && err != nil {
			continue
		}
		if err != nil {
			log.Error("Failed check exist of id:", err.Error())
			continue
		}

		err = s.searchEngine.UpdateArticle(*article)
		if err != nil {
			continue
		}
		msg.Ack(true)
	}
}
