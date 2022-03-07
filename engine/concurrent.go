package engine

import (
	"go-crawler/common"
	"go-crawler/scheduler"
)

type Concurrent struct {
	Worker int
}

func (c Concurrent) Run(seeds ...common.Request) {
	master := make(chan common.Request)
	result := make(chan common.ParseResult)
	s := scheduler.Simple{Master: master}

	for _, seed := range seeds {
		s.Submit(seed)
	}

	for i := 0; i < c.Worker; i++ {
		createWorker(master, result)
	}

	for {
		rs := <-result
		for _, seed := range rs.Requests {
			s.Submit(seed)
		}
	}
}

func createWorker(master chan common.Request, result chan common.ParseResult) {
	go func() {
		for {
			req := <-master
			parseResult, err := work(req)
			if err != nil {
				continue
			}
			result <- *parseResult
		}
	}()
}
