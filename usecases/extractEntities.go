package usecases

import (
	"bytes"
	"encoding/json"
	"github.com/pressus/models/presenters"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type NERReq struct {
	Data string `json:"data"`
}

func (s *service) ExtractEntities(id string) ([]presenters.Entity, error) {

	article, err := s.searchEngine.GetArticleByID(id)
	if err != nil {
		return nil, err
	}

	reqObj := &NERReq{Data: article.Body}
	reqData, _ := json.Marshal(reqObj)
	reqBody := reqData
	client := http.Client{}
	url := "http://" + s.env.Config.NERAPI.IP + ":" + s.env.Config.NERAPI.Port + "/ner"
	req, err := http.NewRequest(http.MethodPost,
		url, bytes.NewBuffer(reqBody))

	if err != nil {
		log.Error("Failed get ner data: ", err.Error())
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed get ner data: ", id, err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Error("Failed get ner data: ", resp.StatusCode)
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Failed read body of get by id response")
		return nil, err
	}

	respObj := []presenters.Entity{}

	err = json.Unmarshal(bodyBytes, &respObj)
	if err != nil {
		log.Error("Failed unmarshal ner response", err.Error())
		return nil, err
	}

	return respObj, nil

}
