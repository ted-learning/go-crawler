package main

import (
	"go-crawler/engine"
	"go-crawler/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://nba.stats.qq.com/team/list.htm",
		ParserFunc: parser.ParseIndex,
	})
}
