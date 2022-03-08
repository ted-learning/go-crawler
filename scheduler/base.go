package scheduler

import "go-crawler/common"

type Scheduler interface {
	Submit(request common.Request)
	GetWorkerChan() chan common.Request
	Run()
	Notification
}

type Notification interface {
	Notify(workerChan chan common.Request)
}
