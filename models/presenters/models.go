package presenters

type ResponseStruct struct {
	Data   []string `json:"data"`
	Status string   `json:"status"`
}

type GetFlowsResp struct {
	Data   []FlowObj `json:"data"`
	Status string    `json:"status"`
}

type FlowObj struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
