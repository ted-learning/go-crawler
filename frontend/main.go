package main

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"go-crawler/common"
	persist "go-crawler/distribute/persist/client"
	"go-crawler/engine"
	"go-crawler/frontend/controller"
	"go-crawler/parser"
	"go-crawler/scheduler"
	"net/http"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	e := engine.Concurrent{
		Worker: 200,
		//Scheduler: &scheduler.Simple{},
		Scheduler: &scheduler.Queue{},
		SaverChan: persist.DataSaver(":1234"),
	}
	jobPool := e.Run()

	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/team/search", controller.CreateSearchTeamResultHandler(
		"frontend/view/temp/teams.html", client,
	))
	http.HandleFunc("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		jobPool.Submit(common.Request{
			Url:        "https://nba.stats.qq.com/team/list.htm",
			ParserFunc: parser.ParseIndex,
		})

		response := struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		}{
			Msg:  "ok",
			Code: 0,
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(response)
		return
	})
	err = http.ListenAndServe(":8888", nil)
	common.PanicErr(err)
}
