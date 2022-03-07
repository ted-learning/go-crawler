package engine

import (
	"encoding/json"
	"go-crawler/common"
	"go-crawler/fetcher"
	"log"
)

type Simple struct{}

func (Simple) Run(seeds ...common.Request) {
	var requests []common.Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		result, err := work(req)
		if err != nil {
			log.Printf("Fetcher: error fetching URL %s: %v\n", req.Url, err)
			continue
		}
		requests = append(requests, result.Requests...)
	}
}

func work(req common.Request) (*common.ParseResult, error) {
	bytes, err := fetcher.Fetcher(req.Url)
	if err != nil {
		return nil, err
	}
	result := req.ParserFunc(bytes, req.Context)
	marshal, err := json.Marshal(result.Result)
	if err != nil {
		return nil, err
	}

	log.Printf("Fetched result: %s\n", marshal)
	return &result, nil
}
