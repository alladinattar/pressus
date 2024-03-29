package presenters

import "time"

type ResponseStruct struct {
	Data   []string `json:"data"`
	Status string   `json:"status"`
}

type ResponseStructInterface struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

type ResponseStructWithArticles struct {
	Data   []ArticleObj `json:"data"`
	Status string       `json:"status"`
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
	ID      string    `json:"id,omitempty"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Authors string    `json:"authors"`
	Link    string    `json:"link"`
	Views   int       `json:"views"`
	Body    string    `json:"body"`
	Flow    string    `json:"flow"`
}

type Entity struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
