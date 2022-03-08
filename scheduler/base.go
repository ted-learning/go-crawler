package scheduler

import "go-crawler/common"

type Scheduler interface {
	Submit(request common.Request)
	ConfigMaster(master chan common.Request)
}
