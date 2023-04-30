package search

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type SearchRespBodyWithTitle struct {
	Hits struct {
		Hits []struct {
			Fields struct {
				Title []string `json:"title"`
			} `json:"fields"`
		} `json:"hits"`
	} `json:"hits"`
}

func (s *engineRepo) FindByFlowAndAuthor(flow, author string) ([]string, error) {
	requestBody := SearchByFlowAndAuthorRequest(flow, author)
	client := http.Client{}
	url := "http://" + s.env.Config.SearchEngine.Ip + ":" + s.env.Config.SearchEngine.Port + "/articles/_search"
	checkArticleReq, err := http.NewRequest(http.MethodGet,
		url,
		bytes.NewBuffer(requestBody))
	checkArticleReq.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error("Failed search by flow and author: ", err.Error())
		return nil, err
	}

	resp, err := client.Do(checkArticleReq)
	if err != nil {
		log.Error("Failed make search request: ", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed make search request: ", resp.StatusCode)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed read body of search response")
		return nil, err
	}

	respObj := &SearchRespBodyWithTitle{}

	err = json.Unmarshal(bodyBytes, respObj)
	if err != nil {
		log.Error("Failed unmarshal search response", err.Error())
		return nil, err
	}

	return collectResults(respObj), nil
}

func collectResults(result *SearchRespBodyWithTitle) []string {
	var items []string
	for _, item := range result.Hits.Hits {
		items = append(items, item.Fields.Title[0])
	}
	return items
}
