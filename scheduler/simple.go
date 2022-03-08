package scheduler

import (
	"go-crawler/common"
)

type Simple struct {
	Master chan common.Request
}

func (s *Simple) Submit(request common.Request) {
	go func() {
		s.Master <- request
	}()
}

func (s *Simple) ConfigMaster(master chan common.Request) {
	s.Master = master
}
