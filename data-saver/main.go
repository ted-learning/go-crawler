package main

import (
	"data-saver/common"
	"data-saver/persist"
	"data-saver/rpc"
	"fmt"
	"github.com/olivere/elastic/v7"
	"net/http"
	"os"
)

func main() {
	esHost := os.Getenv("ElasticHost")
	if esHost == "" {
		esHost = "127.0.0.1"
	}
	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%s:9200", esHost)), elastic.SetSniff(false))
	common.PanicErr(err)

	service := &persist.DataSaverRpcService{
		Client: client,
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			common.PanicErr(err)
		}
	})
	err = http.ListenAndServe(":12341", nil)
	if err != nil {
		common.PanicErr(err)
	}

	rpcsupport.ServeRpc(":1234", service)
}
