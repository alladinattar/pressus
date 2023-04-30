package search

import "time"

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

func SearchByFlowAndDateRequest(flow string, from, until time.Time) []byte {
	requestBody := []byte(`
{
    "query": {
        "bool": {
            "should": [
                {
                    "match": {
                        "flow": "` + flow + `"
                    }
                },
                {
                    "range": {
                        "date": {
                            "gte": "` + from.Format("2006-01-02") + `",
                            "lte":"` + until.Format("2006-01-02") + `"
                        }
                    }
                }
            ]
        }
    },
    "fields": ["title"],
    "_source": false
}
`)
	return requestBody
}
