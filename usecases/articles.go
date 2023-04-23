package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

type Articles struct {
	mu       sync.Mutex
	Articles []presenters.ArticleLink
}

func (s *service) GetArticlesByFlow(flow string) ([]presenters.ArticleLink, error) {
	articles, err := s.extractArticles(flow)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *service) extractArticles(flow string) ([]presenters.ArticleLink, error) {
	pages := make(chan int)
	var articles Articles
	go s.checkPages(flow, pages)
	var wg sync.WaitGroup
	for page := range pages {
		wg.Add(1)
		go s.parseArticles(&wg, &articles, flow, strconv.Itoa(page))
	}
	wg.Wait()
	return articles.Articles, nil
}

func (s *service) checkPages(flow string, pages chan<- int) {
	client := fiber.Client{
		UserAgent: s.env.Config.Parser.UserAgent,
	}
	for i := 1; ; i++ {
		requestString := fmt.Sprintf("%s/%s/%s/page/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow, strconv.Itoa(i))
		var resp []byte
		statusCode, _, err := client.Get(requestString).Timeout(time.Second*5).Get(resp, requestString)
		if err != nil {
			log.Error(err)
		}
		if statusCode == fiber.StatusNotFound {
			close(pages)
			return
		} else if statusCode == fiber.StatusOK {
			log.Println("Find page: ", i)
			pages <- i
		}
	}
}

func (s *service) parseArticles(wg *sync.WaitGroup, articles *Articles, flow, page string) error {
	defer wg.Done()
	client := fiber.Client{
		UserAgent: s.env.Config.Parser.UserAgent,
	}

	requestString := fmt.Sprintf("%s/%s/%s/page/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows", flow, page)
	var resp []byte
	_, body, err := client.Get(requestString).Timeout(time.Second*5).Get(resp, requestString)
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return err
	}

	doc.Find(".link--MuU14").Each(func(i int, s *goquery.Selection) {
		link, _ := s.Attr("href")
		article := presenters.ArticleLink{
			Title: s.Text(),
			Link:  link,
		}
		articles.mu.Lock()
		defer articles.mu.Unlock()
		articles.Articles = append(articles.Articles, article)
	})
	return nil
}
