package main

import (
	"github.com/olivere/elastic/v7"
	"go-crawler/common"
	"go-crawler/frontend/controller"
	"net/http"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/team/search", controller.CreateSearchTeamResultHandler(
		"frontend/view/temp/teams.html", client,
	))
	err = http.ListenAndServe(":8888", nil)
	common.PanicErr(err)
}
