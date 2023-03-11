package usecases

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"github.com/pressus/models/presenters"
)

func (s *service) GetFlows() ([]presenters.FlowObj, error) {
	client := fiber.Client{
		UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	}
	requestString := fmt.Sprintf("%s/%s/", s.GetEnv().Config.Parser.DefaultRoute, "flows")
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

func (s *service) getFlows(data []byte) ([]presenters.FlowObj, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	var flows []presenters.FlowObj
	doc.Find(".item--QHDrx").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		link, _ := s.Attr("href")
		flow := presenters.FlowObj{
			Name: name,
			Link: link,
		}
		flows = append(flows, flow)
	})
	return flows, nil
}
