package search

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type CountRequestBody struct {
	Query struct {
		Match struct {
			ID string `json:"id"`
		} `json:"match"`
	} `json:"query"`
}

type CountResponseBody struct {
	Count int `json:"count"`
}

func (s *engineRepo) IsArticleExist(id string) (bool, error) {
	requestObj := CountRequestBody{}
	requestObj.Query.Match.ID = id
	body, _ := json.MarshalIndent(requestObj, "", "  ")
	client := http.Client{}
	url := "http://" + s.env.Config.SearchEngine.Ip + ":" + s.env.Config.SearchEngine.Port + "/articles/_count"
	checkArticleReq, err := http.NewRequest(http.MethodGet,
		url,
		bytes.NewBuffer(body))
	checkArticleReq.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error("Failed check existence: ", err.Error())
		return true, err
	}

	resp, err := client.Do(checkArticleReq)
	if err != nil {
		log.Error("Failed get count of articles with id: ", id, err.Error())
		return true, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed check exists id article: ", resp.StatusCode)
		return true, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed read body of count response")
		return true, err
	}

	respObj := &CountResponseBody{}

	err = json.Unmarshal(bodyBytes, respObj)
	if err != nil {
		log.Error("Failed unmarshal count response", err.Error())
		return true, err
	}

	if respObj.Count != 1 {
		return false, nil
	}

	return true, nil
}
