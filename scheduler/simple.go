package scheduler

import (
	"go-crawler/common"
)

type Simple struct {
	Master chan common.Request
}

func (s *Simple) WorkerReady(_ chan common.Request) {}

func (s *Simple) Run() {
	s.Master = make(chan common.Request)
}

func (s *Simple) GetWorkerChan() chan common.Request {
	return s.Master
}

func (s *Simple) Submit(request common.Request) {
	go func() {
		s.Master <- request
	}()
}
