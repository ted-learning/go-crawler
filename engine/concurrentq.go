package engine

import (
	"go-crawler/common"
	"go-crawler/scheduler"
)

type ConcurrentQ struct {
	Worker    int
	Scheduler *scheduler.Queue
}

func (c ConcurrentQ) Run(seeds ...common.Request) {
	c.Scheduler.Run()
	result := make(chan common.ParseResult)
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}

	for i := 0; i < c.Worker; i++ {
		worker := createWorkerQ(result)
		c.Scheduler.ConfigWorker(worker)
	}

	for {
		rs := <-result
		for _, seed := range rs.Requests {
			c.Scheduler.Submit(seed)
		}
	}
}

func createWorkerQ(result chan common.ParseResult) common.Worker {
	worker := make(chan common.Request)
	go func() {
		for {
			req := <-worker
			parseResult, err := work(req)
			if err != nil {
				continue
			}
			result <- *parseResult
		}
	}()
	return worker
}
