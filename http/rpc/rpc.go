package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/thoohv5/util/http"
)

const (
	Version   = "2.0"
	DefaultID = 1
	TimeOut   = 3
)

type Svr interface {
	GetConfig(name string) *Config
	ModuleName() string
	ServiceName() string
}

// Config rpc
type Config struct {
	Host    string `toml:"host"`
	TimeOut uint32 `toml:"timeout"`
}

type ReqArgs struct {
	JsonRpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      uint8       `json:"id"`
}

type Resp struct {
	JsonRpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      uint8       `json:"id"`
	Error   interface{} `json:"error"`
}

// Call 支持回调
func Call(ctx context.Context, rpc Svr, args *ReqArgs, result interface{}) error {

	moduleName := rpc.ModuleName()
	rpcConfig := rpc.GetConfig(moduleName)
	baseUrl := rpcConfig.Host
	timeout := rpcConfig.TimeOut

	if timeout == 0 {
		timeout = TimeOut
	}

	url := fmt.Sprintf("%s/%s", baseUrl, rpc.ServiceName())
	args.JsonRpc = Version
	args.ID = DefaultID

	// 返回结果
	rpcResp := Resp{}
	if nil != result {
		rpcResp.Result = &result
	}

	// 请求
	if err := http.Post(ctx, url, nil, &rpcResp, http.WithPreDeal(func(r *http.Parameter) error {
		data, err := json.Marshal(args)
		if nil != err {
			return fmt.Errorf("json Marshal err:%w", err)
		}
		r.SetBody(bytes.NewBuffer(data))
		return nil
	}), http.WithHeader(map[string]string{
		"Content-Type": "application/json",
	}), http.WithTimeout(time.Duration(timeout)*time.Second)); nil != err {
		return fmt.Errorf("post err: %w", err)
	}

	// JsonRPC 报错
	if nil != rpcResp.Error {
		rpcErrorByte, _ := json.Marshal(rpcResp.Error)
		return fmt.Errorf("rpc Call error: url:%s, rpcErr:%s", url, string(rpcErrorByte))
	}

	return nil
}
