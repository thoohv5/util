package example

import (
	"github.com/thoohv5/util/http/rpc"
)

type JsonRpc struct {
}

func (jr JsonRpc) GetConfig(name string) *rpc.Config {
	return &rpc.Config{
		Host:    "",
		TimeOut: 20,
	}
}
