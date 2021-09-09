package im

import "github.com/thoohv5/util/http/rpc/example"

// IM服务
type SvrRpc struct {
	example.JsonRpc
}

func (svr SvrRpc) ModuleName() string {
	return "im"
}
