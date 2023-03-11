package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

func (s *service) GetArticlesByFlow(flow string) ([]string, error) {
	articles, err := s.extractArticles(flow)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *service) extractArticles(flow string) ([]string, error) {
	pages := make(chan int)
	var articles []string
	go s.checkPages(flow, pages)
	var wg sync.WaitGroup
	for page := range pages {
		wg.Add(1)
		go s.parseArticles(&wg, &articles, flow, strconv.Itoa(page))
	}
	wg.Wait()
	return articles, nil
}

func (s *service) checkPages(flow string, pages chan<- int) {
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	for i := 1; ; i++ {
		requestString := fmt.Sprintf("%s/%s/%s/page/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow, strconv.Itoa(i))
		var resp []byte
		statusCode, _, err := client.Head(requestString).Get(resp, requestString)
		if err != nil {
			log.Error(err)
		}
		if statusCode == fiber.StatusNotFound {
			close(pages)
			return
		} else if statusCode == fiber.StatusOK {
			pages <- i
		}
	}
}

func (s *service) parseArticles(wg *sync.WaitGroup, articles *[]string, flow, page string) error {
	defer wg.Done()
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}

	requestString := fmt.Sprintf("%s/%s/%s/page/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow, page)
	var resp []byte
	_, body, err := client.Get(requestString).Get(resp, requestString)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return err
	}
	doc.Find(".title--zzk3s").Each(func(i int, s *goquery.Selection) {
		*articles = append(*articles, s.Text())
	})
	return nil
}
