package main

import (
	"go-crawler/common"
	"go-crawler/engine"
	"go-crawler/parser"
	"go-crawler/persist"
	"go-crawler/scheduler"
)

func main() {
	e := engine.Concurrent{
		Worker:    200,
		Scheduler: &scheduler.Simple{},
		//Scheduler: &scheduler.Queue{},
		SaverChan: persist.DataSaver(),
	}

	e.Run(common.Request{
		Url:        "https://nba.stats.qq.com/team/list.htm",
		ParserFunc: parser.ParseIndex,
	})
}
