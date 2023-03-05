package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func (s *service) GetArticlesByFlow(flow string) ([]string, error) {
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	requestString := fmt.Sprintf("%s/%s/%s", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow)
	var resp []byte
	statusCode, body, err := client.Get(requestString).Get(resp, requestString)
	if err != nil {
		panic(err)
	}
	if statusCode != 200 {
		panic("Not 200 status code")
	}
	articles, err := s.extractArticles(body)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *service) extractArticles(data []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var articles []string
	doc.Find(".title--zzk3s").Each(func(i int, s *goquery.Selection) {
		articles = append(articles, s.Text())
	})
	return articles, nil

}
