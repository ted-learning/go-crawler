package main

import (
	"github.com/olivere/elastic/v7"
	"go-crawler/common"
	"go-crawler/distribute/persist"
	rpcsupport "go-crawler/distribute/rpc"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	service := &persist.DataSaverRpcService{
		Client: client,
	}
	rpcsupport.ServeRpc(":1234", service)
}
