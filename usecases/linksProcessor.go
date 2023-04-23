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
	"strings"
	"time"
)

func (s *service) ProcessLinks() {
	msgs := make(chan amqp091.Delivery)
	go s.repo.GetTasks(msgs)

	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Mobile Safari/537.36",
	}
	for msg := range msgs {
		task := &presenters.ArticleLink{}
		err := json.Unmarshal(msg.Body, &task)
		if err != nil {
			log.Error("Failed unmarshall task: ", err.Error())
		}
		log.Info("Received from tasks: ", task.Title)
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

		article := &presenters.ArticleObj{}

		var authors []string
		doc.Find(".article-author__name").Each(func(i int, sel *goquery.Selection) {
			author := strings.TrimSpace(sel.Text())
			authors = append(authors, author)
		})
		article.Authors = authors

		doc.Find("._27USv").Each(func(i int, sel *goquery.Selection) {
			timeAttr, _ := sel.Attr("datetime")
			if timeAttr != "" {
				date, err := time.Parse("02.01.06", sel.Text())
				if err != nil {
					log.Error("Failed parse date:", err.Error())
				}
				article.Date = date
			}
		})

		doc.Find(".article-header__title").Each(func(i int, sel *goquery.Selection) {
			article.Body = sel.Text()
			s.repo.PushArticleToResults(article)
		})
		msg.Ack(true)
	}
}

func (s *service) ProcessLinksFromResultQueue() {
	articles := make(chan amqp091.Delivery)
	go s.repo.GetResults(articles)

	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	for msg := range articles {
		task := &presenters.ArticleLink{}
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
			s.repo.PushArticleToRusults(sel.Text())
		})
		msg.Ack(true)
	}
}
