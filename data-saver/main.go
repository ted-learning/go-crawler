package main

import (
	"data-saver/common"
	"data-saver/persist"
	"data-saver/rpc"
	"github.com/olivere/elastic/v7"
)

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	common.PanicErr(err)

	service := &persist.DataSaverRpcService{
		Client: client,
	}
	rpcsupport.ServeRpc(":1234", service)
}
