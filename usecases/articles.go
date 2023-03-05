package usecases

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (s *service) GetArticlesByFlow(flow string) ([]string, error) {
	articles, err := s.extractArticles(flow)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *service) extractArticles(flow string) ([]string, error) {
	var articles []string
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	for i := 1; ; i++ {
		requestString := fmt.Sprintf("%s/%s/%s/page/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow, strconv.Itoa(i))
		var resp []byte
		statusCode, body, err := client.Get(requestString).Get(resp, requestString)
		if err != nil {
			panic(err)
		}
		if statusCode == 404 {
			return nil, errors.New("invalid flow")
		} else if statusCode != 200 {
			return nil, errors.New("unknown error")
		}

		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}

		doc.Find(".title--zzk3s").Each(func(i int, s *goquery.Selection) {
			articles = append(articles, s.Text())
		})

		if s.isLatestPage(doc, string(i)) {
			return articles, nil
		}
	}

	return articles, nil
}

func (s *service) isLatestPage(doc *goquery.Document, currentPage string) bool {
	var pageNumbers []string
	doc.Find(".link--hAARL").Each(func(i int, selection *goquery.Selection) {
		pageNumbers = append(pageNumbers, selection.Text())
	})

	if len(pageNumbers) == 0 {
		return true
	}

	if pageNumbers[len(pageNumbers)-1] == currentPage {
		return true
	}
	return false
}
