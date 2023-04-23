package search

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func (s *engineRepo) SaveArticle(obj presenters.ArticleObj) error {
	hashID := md5.Sum([]byte(obj.Title + obj.Date.String()))
	body, _ := json.Marshal(obj)
	client := http.Client{}
	url := "http://" + s.env.Config.SearchEngine.Ip + ":" + s.env.Config.SearchEngine.Port + "/articles/_doc/" + fmt.Sprintf("%x", hashID)
	addArticleReq, err := http.NewRequest(http.MethodPut,
		url,
		bytes.NewBuffer(body))
	addArticleReq.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Error("Failed add article: ", err.Error())
		return err
	}

	resp, err := client.Do(addArticleReq)
	if err != nil {
		log.Error("Failed add article to elastic: ", err.Error())
		return err
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		log.Error("Failed add article: ", resp.StatusCode)
		return err
	}
	return nil
}
