package app

import (
	"bytes"
	"github.com/pressus/config"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func initElastic(env config.Env) {
	url := "http://" + env.Config.SearchEngine.Ip + ":" + env.Config.SearchEngine.Port + "/articles"
	settings := []byte(`{
	"settings" : {
		"number_of_shards" : 1,
		"number_of_replicas" : 0
	},
	"mappings": {
		"properties": {
				"id": {
					"type": "text"
				},
				"title": {
					"type": "text"
				},
				"date": {
					"type": "date"
				},
				"authors": {
					"type":"text"
				},
				"link": {
					"type":"text"
				}
			}
		}
}`)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(settings))
	if err != nil {
		log.Fatal("Failed init elastic index:", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Failed init elastic index:", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusBadRequest {
			log.Info("Index already exists:", resp.Status)
		} else {
			log.Fatal("Failed init elastic index: ", resp.Status)
		}
	}
}
