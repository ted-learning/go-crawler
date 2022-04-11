package main

import (
	"data-saver/common"
	"data-saver/persist"
	"data-saver/rpc"
	"fmt"
	"github.com/olivere/elastic/v7"
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
	rpcsupport.ServeRpc(":1234", service)
}
