package main

import (
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
	"os"
	"portal/common"
	"portal/controller"
)

func main() {
	esHost := os.Getenv("ElasticHost")
	if esHost == "" {
		esHost = "127.0.0.1"
	}
	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:9200", esHost)), elastic.SetSniff(false))
	common.PanicErr(err)

	http.Handle("/", http.FileServer(http.Dir("view")))
	http.Handle("/team/search", controller.CreateSearchTeamResultHandler(
		"view/temp/teams.html", client,
	))

	engineHost := os.Getenv("EngineHost")
	if engineHost == "" {
		engineHost = "127.0.0.1"
	}
	http.HandleFunc("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		_, err := http.Post(fmt.Sprintf("http://%s:7500/refresh", engineHost), "application/json", nil)
		if err != nil {
			log.Printf(err.Error())
			return
		}
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
