package server

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	defaultSeverMan = NewServerMan()
)

const (
	Timeout = 3 * time.Second
)

type man struct {
	svrs  []Server
	async bool
	wait  bool
}

type Option interface {
	apply(*man)
}

type optionFunc func(*man)

func (f optionFunc) apply(o *man) {
	f(o)
}

func WithAsync(async bool) Option {
	return optionFunc(func(o *man) { o.async = async })
}

func WithWait(wait bool) Option {
	return optionFunc(func(o *man) { o.wait = wait })
}

func NewServerMan(opts ...Option) *man {
	m := &man{
		async: true,
		wait:  false,
	}
	for _, o := range opts {
		o.apply(m)
	}

	m.svrs = make([]Server, 0, 1)
	return m
}

func (m *man) RegisterServer(server Server) {
	m.svrs = append(m.svrs, server)
}

func (m *man) Start() {
	wg := sync.WaitGroup{}
	wg.Add(len(m.svrs))
	for _, svr := range m.svrs {
		if m.async {
			go func(s Server) {
				defer wg.Done()
				s.Serve()
			}(svr)
		} else {
			func(s Server) {
				defer wg.Done()
				s.Serve()
			}(svr)
		}
	}
	wg.Wait()
	if !m.wait {
		go handleSysSignal()
	} else {
		handleSysSignal()
	}
}

func (m *man) Stop(ctx context.Context) {
	for _, svr := range m.svrs {
		svr.Stop(ctx)
	}
}

func RegisterServers(servers ...Server) {
	for _, server := range servers {
		defaultSeverMan.RegisterServer(server)
	}
}

func Start(opts ...Option) {
	for _, o := range opts {
		o.apply(defaultSeverMan)
	}
	defaultSeverMan.Start()
}

func Stop(ctx context.Context) {
	defaultSeverMan.Stop(ctx)
}

func handleSysSignal() {
	sChan := make(chan os.Signal, 1)
	signal.Notify(sChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	for {
		s := <-sChan
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			ctx, cancel := context.WithTimeout(context.Background(), Timeout)
			defer func() {
				cancel()
			}()
			Stop(ctx)
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}

	}
}
