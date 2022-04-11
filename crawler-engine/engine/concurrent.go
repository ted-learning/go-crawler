package engine

import (
	"crawler-engine/common"
	"crawler-engine/scheduler"
)

type Concurrent struct {
	Worker    int
	Scheduler scheduler.Scheduler
	SaverChan chan interface{}
	out       chan common.ParseResult
}

func (c *Concurrent) Run() scheduler.Scheduler {
	c.Scheduler.Run()
	c.out = make(chan common.ParseResult)
	for i := 0; i < c.Worker; i++ {
		c.createWorker()
	}

	go func() {
		for {
			rs := <-c.out
			go func() { c.SaverChan <- rs.Result }()
			for _, seed := range rs.Requests {
				c.Scheduler.Submit(seed)
			}
		}
	}()
	return c.Scheduler
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
