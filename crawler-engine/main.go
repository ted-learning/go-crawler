package main

import (
	"crawler-engine/common"
	"crawler-engine/distribute/persist/client"
	"crawler-engine/engine"
	"crawler-engine/parser"
	"crawler-engine/scheduler"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	dataSaverHost := os.Getenv("DataSaverHost")
	if dataSaverHost == "" {
		dataSaverHost = "127.0.0.1"
	}
	e := engine.Concurrent{
		Worker:    200,
		Scheduler: &scheduler.Simple{},
		//Scheduler: &scheduler.Queue{},
		SaverChan: client.DataSaver(fmt.Sprintf("%s:1234", dataSaverHost)),
	}

	jobPool := e.Run()
	http.HandleFunc("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		jobPool.Submit(
			common.Request{
				Url:        "https://nba.stats.qq.com/team/list.htm",
				ParserFunc: parser.ParseIndex,
			},
		)

		response := struct {
			Msg  string `json:"msg"`
			Code int    `json:"code"`
		}{
			Msg:  "ok",
			Code: 0,
		}
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(response)
		if err != nil {
			common.PanicErr(err)
		}
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			common.PanicErr(err)
		}
	})

	err := http.ListenAndServe(":7500", nil)
	common.PanicErr(err)

	for {
		time.Sleep(time.Duration(1000))
	}
}
