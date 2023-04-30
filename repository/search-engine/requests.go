package search

func SearchByFlowAndAuthorRequest(flow, author string) []byte {
	requestBody := []byte(`
{
    "query": {
        "bool": {
            "must": [
                {
                    "match": {
                        "authors": "` + author + `"
                    }
                },
                {
                    "match": {
                        "flow": "` + flow + `"
                    }
                }
            ]
        }
    },
   "fields":["title"],
    "_source": false
}
`)
	return requestBody
}
