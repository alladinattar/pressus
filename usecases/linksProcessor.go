package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
)

func (s *service) ProcessLinks() {
	tasks := make(chan presenters.ArticleObj)
	go s.repo.GetTasks(tasks)

	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	for task := range tasks {
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
	}
}
