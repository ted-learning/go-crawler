package main

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"net/http"
	"portal/common"
	"portal/controller"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/team/search", controller.CreateSearchTeamResultHandler(
		"frontend/view/temp/teams.html", client,
	))
	http.HandleFunc("/refresh", func(writer http.ResponseWriter, request *http.Request) {
		_, err := http.Get("http://localhost:9200/refresh")
		if err != nil {
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
