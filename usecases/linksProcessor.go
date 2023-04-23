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
	go s.repoTasks.GetTasks(msgs)
	client := fiber.Client{
		UserAgent: s.env.Config.Parser.UserAgent,
	}

	for msg := range msgs {
		task := &presenters.ArticleLink{}
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

		doc.Find(".article-body").Each(func(i int, sel *goquery.Selection) {
			if sel.
			article.Body = sel.Text()
			log.Printf(sel.Text())
		})
		article.Flow = strings.Replace(task.Link, "/", "", -1)
		article.Title = task.Title
		article.Link = task.Link
		s.repoTasks.PushArticleToResults(article)
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

		err = s.searchEngine.SaveArticle(*article)
		if err != nil {
			continue
		}
		msg.Ack(true)
	}
}
