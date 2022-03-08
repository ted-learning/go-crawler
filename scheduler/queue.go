package scheduler

import (
	"go-crawler/common"
)

type Queue struct {
	requests chan common.Request
	workers  chan common.Worker
}

func (q *Queue) ConfigMaster(_ chan common.Request) {
	panic("implement me")
}

func (q *Queue) Submit(request common.Request) {
	q.requests <- request
}

func (q *Queue) ConfigWorker(worker common.Worker) {
	q.workers <- worker
}

func (q *Queue) Run() {
	var workerQ []common.Worker
	var requestQ []common.Request
	for {
		var activeRequest common.Request
		var activeWorker common.Worker
		if len(workerQ) > 0 && len(requestQ) > 0 {
			activeRequest = requestQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case request := <-q.requests:
			requestQ = append(requestQ, request)
		case worker := <-q.workers:
			workerQ = append(workerQ, worker)
		case activeWorker <- activeRequest:
			requestQ = requestQ[1:]
			workerQ = workerQ[1:]
		}
	}
}
