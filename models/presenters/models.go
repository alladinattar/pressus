package presenters

import "time"

type ResponseStruct struct {
	Data   []string `json:"data"`
	Status string   `json:"status"`
}

type StatusResponseStruct struct {
	Status string `json:"status"`
}

type GetFlowsResp struct {
	Data   []FlowObj `json:"data"`
	Status string    `json:"status"`
}

type FlowObj struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type ArticleLink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type ArticleObj struct {
	Title  string    `json:"title"`
	Date   time.Time `json:"date"`
	Author string    `json:"author"`
	Link   string    `json:"link"`
}
