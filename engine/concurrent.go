package engine

import (
	"go-crawler/common"
	"go-crawler/scheduler"
)

type Concurrent struct {
	Worker    int
	Scheduler scheduler.Scheduler
}

func (c Concurrent) Run(seeds ...common.Request) {
	result := make(chan common.ParseResult)
	c.Scheduler.Run()

	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}

	for i := 0; i < c.Worker; i++ {
		workerChan := c.Scheduler.GetWorkerChan()
		createWorker(workerChan, result)
		c.Scheduler.Notify(workerChan)
	}

	for {
		rs := <-result
		for _, seed := range rs.Requests {
			c.Scheduler.Submit(seed)
		}
	}
}

func createWorker(worker chan common.Request, result chan common.ParseResult) {
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
}
