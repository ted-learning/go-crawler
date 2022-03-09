package scheduler

import "go-crawler/common"

type Scheduler interface {
	Submit(request common.Request)
	GetWorkerChan() chan common.Request
	Run()
	Notify
}

type Notify interface {
	WorkerReady(workerChan chan common.Request)
}
