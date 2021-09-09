package server

import "context"

var (
	svrBox []Server
)

type Server interface {
	Serve() error
	Stop(ctx context.Context) error
}

func RegisterServer(svr Server) {
	svrBox = append(svrBox, svr)
}
