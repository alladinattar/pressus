package search

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

func (s *engineRepo) FindByFlowAndDate(flow string, from, until time.Time) ([]string, error) {
	requestBody := SearchByFlowAndDateRequest(flow, from, until)
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
