package engine

import (
	"go-crawler/common"
	"go-crawler/scheduler"
)

type Concurrent struct {
	Worker    int
	Scheduler scheduler.Scheduler
	SaverChan chan interface{}
	out       chan common.ParseResult
}

func (c *Concurrent) Run(seeds ...common.Request) {
	c.out = make(chan common.ParseResult)
	for i := 0; i < c.Worker; i++ {
		c.createWorker()
	}

	c.Scheduler.Run()
	for _, seed := range seeds {
		c.Scheduler.Submit(seed)
	}

	for {
		rs := <-c.out
		go func() { c.SaverChan <- rs.Result }()
		for _, seed := range rs.Requests {
			c.Scheduler.Submit(seed)
		}
	}
}

func (c *Concurrent) createWorker() {
	workerChan := c.Scheduler.GetWorkerChan()
	go func() {
		for {
			c.Scheduler.WorkerReady(workerChan)
			req := <-workerChan
			parseResult, err := work(req)
			if err != nil {
				continue
			}
			c.out <- *parseResult
		}
	}()
}
