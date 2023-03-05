package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
)

func (s *service) GetFlows() ([]string, error) {
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	requestString := fmt.Sprintf("%s/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows")
	fmt.Printf(requestString)
	var resp []byte
	statusCode, body, err := client.Get(requestString).Get(resp, requestString)
	if err != nil {
		panic(err)
	}
	if statusCode != 200 {
		panic("Not 200 status code")
	}
	flows, err := s.getFlows(body)
	if err != nil {
		return nil, err
	}

	return flows, nil
}

func (s *service) getFlows(data []byte) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var flows []string
	doc.Find(".text--zYMfN").Each(func(i int, s *goquery.Selection) {
		flows = append(flows, s.Text())
	})
	return flows, nil
}
