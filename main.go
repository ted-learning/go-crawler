package main

import (
	"go-crawler/common"
	"go-crawler/engine"
	"go-crawler/parser"
	"go-crawler/scheduler"
	"log"
	"time"
)

func main() {
	s := time.Now()
	engine.Concurrent{
		Worker:    100,
		Scheduler: &scheduler.Simple{},
	}.Run(common.Request{
		Url:        "https://nba.stats.qq.com/team/list.htm",
		ParserFunc: parser.ParseIndex,
	})
	e := time.Now()
	log.Printf("Total crawler cost: %f", e.Sub(s).Seconds())
}
