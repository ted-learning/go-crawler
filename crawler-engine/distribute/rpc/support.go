package rpcsupport

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewClient(host string) (*rpc.Client, error) {
	dial, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(dial), nil
}
