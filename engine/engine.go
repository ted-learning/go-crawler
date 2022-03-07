package engine

import (
	"encoding/json"
	"go-crawler/common"
	"go-crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		bytes, err := fetcher.Fetcher(req.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching URL %s: %v\n", req.Url, err)
			continue
		}
		result := req.ParserFunc(bytes, req.Context)
		marshal, err := json.Marshal(result.Result)
		common.PanicErr(err)

		log.Printf("Fetched result: %s\n", marshal)
		requests = append(requests, result.Requests...)
	}
}
