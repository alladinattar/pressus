package search

import (
	"bytes"
	"encoding/json"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type ArticleResp struct {
	Source presenters.ArticleObj `json:"_source"`
}

func (s *engineRepo) GetArticleByID(id string) (*presenters.ArticleObj, error) {
	client := http.Client{}
	url := "http://" + s.env.Config.SearchEngine.Ip + ":" + s.env.Config.SearchEngine.Port + "/articles/_doc/" + id
	req, err := http.NewRequest(http.MethodGet,
		url, bytes.NewBuffer(nil))

	if err != nil {
		log.Error("Failed get by id: ", err.Error())
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed get by id: ", id, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed get by id: ", resp.StatusCode)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed read body of get by id response")
		return nil, err
	}

	respObj := &ArticleResp{}

	err = json.Unmarshal(bodyBytes, respObj)
	if err != nil {
		log.Error("Failed unmarshal get by id response", err.Error())
		return nil, err
	}

	return &respObj.Source, nil
}
