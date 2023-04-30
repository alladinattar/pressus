package search

import (
	"bytes"
	"encoding/json"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UpdateRequest struct {
	Doc presenters.ArticleObj `json:"doc"`
}

func (s *engineRepo) UpdateArticle(obj presenters.ArticleObj) error {
	reqBody := UpdateRequest{Doc: obj}
	body, _ := json.Marshal(reqBody)
	client := http.Client{}
	url := "http://" + s.env.Config.SearchEngine.Ip + ":" + s.env.Config.SearchEngine.Port + "/articles/_update/" + obj.ID
	addArticleReq, err := http.NewRequest(http.MethodPost,
		url,
		bytes.NewBuffer(body))
	addArticleReq.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error("Failed update article: ", err.Error())
		return err
	}

	resp, err := client.Do(addArticleReq)
	if err != nil {
		log.Error("Failed update article to elastic: ", err.Error())
		return err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		log.Error("Failed update article: ", resp.StatusCode)
		return err
	}
	return nil
}
