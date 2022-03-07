package main

import (
	"go-crawler/engine"
	"go-crawler/parser"
	"log"
	"time"
)

func main() {
	s := time.Now()
	engine.Run(engine.Request{
		Url:        "https://nba.stats.qq.com/team/list.htm",
		ParserFunc: parser.ParseIndex,
	})
	e := time.Now()
	log.Printf("Total crawler cost: %f", e.Sub(s).Seconds())
}
