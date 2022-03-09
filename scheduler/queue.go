package scheduler

import (
	"go-crawler/common"
)

type Queue struct {
	requestChan chan common.Request
	workerChan  chan chan common.Request
}

func (q *Queue) WorkerReady(workerChan chan common.Request) {
	q.workerChan <- workerChan
}

func (q *Queue) GetWorkerChan() chan common.Request {
	return make(chan common.Request)
}

func (q *Queue) Submit(request common.Request) {
	go func() { q.requestChan <- request }()
}

func (q *Queue) Run() {
	q.requestChan = make(chan common.Request)
	q.workerChan = make(chan chan common.Request)
	go func() {
		var workerQ []chan common.Request
		var requestQ []common.Request
		for {
			var activeRequest common.Request
			var activeWorker chan common.Request
			if len(workerQ) > 0 && len(requestQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case request := <-q.requestChan:
				requestQ = append(requestQ, request)
			case worker := <-q.workerChan:
				workerQ = append(workerQ, worker)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()

}
