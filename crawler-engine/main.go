package main

import (
	"crawler-engine/common"
	"crawler-engine/distribute/persist/client"
	"crawler-engine/engine"
	"crawler-engine/parser"
	"crawler-engine/scheduler"
	"time"
)

func main() {
	e := engine.Concurrent{
		Worker: 200,
		//Scheduler: &scheduler.Simple{},
		Scheduler: &scheduler.Queue{},
		SaverChan: client.DataSaver(":1234"),
	}

	jobPool := e.Run()
	jobPool.Submit(
		common.Request{
			Url:        "https://nba.stats.qq.com/team/list.htm",
			ParserFunc: parser.ParseIndex,
		},
	)

	for {
		time.Sleep(time.Duration(1000))
	}
}
