package main

import (
	"go-crawler/common"
	"go-crawler/engine"
	"go-crawler/parser"
	"go-crawler/persist"
	"go-crawler/scheduler"
)

func main() {
	saverChan := persist.DataSaver()
	e := engine.Concurrent{
		Worker: 100,
		//Scheduler: &scheduler.Simple{},
		Scheduler: &scheduler.Queue{},
		SaverChan: saverChan,
	}

	e.Run(common.Request{
		Url:        "https://nba.stats.qq.com/team/list.htm",
		ParserFunc: parser.ParseIndex,
	})
}
