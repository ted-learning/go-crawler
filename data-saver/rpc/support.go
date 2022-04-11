package rpcsupport

import (
	"data-saver/common"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServeRpc(host string, service interface{}) {
	err := rpc.Register(service)
	common.PanicErr(err)

	listen, err := net.Listen("tcp", host)
	common.PanicErr(err)

	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(accept)
	}
}
