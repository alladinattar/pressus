package usecases

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Articles struct {
	mu       sync.Mutex
	Articles []presenters.ArticleObj
}

func (s *service) GetArticlesByFlow(flow string) ([]presenters.ArticleObj, error) {
	articles, err := s.extractArticles(flow)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *service) extractArticles(flow string) ([]presenters.ArticleObj, error) {
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

	doc.Find(".header--Oqixx").Each(func(i int, sel *goquery.Selection) {
		article := presenters.ArticleObj{}

		timeAttr, exist := sel.Find(".date--80kR0").First().Attr("datetime")
		if exist {
			date, err := time.Parse("2006-01-02", timeAttr)
			if err != nil {
				log.Error("Failed parse date:", err.Error())
			}
			article.Date = date
		}

		title := sel.Find(".title--ZfdMh").First().Text()
		article.Title = title

		hashID := md5.Sum([]byte(article.Title + article.Date.String()))
		article.ID = fmt.Sprintf("%x", hashID)

		if exist, _ := s.searchEngine.IsArticleExist(article.ID); exist {
			log.Info("Article exists: ", article.ID)
			return
		}

		link, exist := sel.Find(".link--MuU14").First().Attr("href")
		if exist {
			article.Link = link
		}

		authorName := sel.Find(".name--Q6uU8").First().Text()

		article.Authors = authorName

		var viewsCount int
		isThousands := strings.Index(sel.Find(".counter--N9V5a").First().Text(), "K")
		if isThousands == -1 {
			viewsCount, _ = strconv.Atoi(sel.Find(".counter--N9V5a").First().Text())
		} else {
			viewsCount, _ = strconv.Atoi(strings.Replace(sel.Find(".counter--N9V5a").First().Text(), "K", "", -1))
			viewsCount = viewsCount * 1000
		}

		article.Views = viewsCount

		article.Flow = flow
		s.searchEngine.SaveArticle(article)
		articles.mu.Lock()
		defer articles.mu.Unlock()
		articles.Articles = append(articles.Articles, article)
	})
	return nil
}
