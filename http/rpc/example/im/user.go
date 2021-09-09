package im

import (
	"context"

	"github.com/thoohv5/util/http/rpc"
)

type UserService interface {
	GetUserInfo(ctx context.Context, req *GetUserInfoReq, resp *GetUserInfoResp) error
}

type userService struct {
	SvrRpc
}

// 创建
func NewUserService() UserService {
	return &userService{}
}

const (
	UserServiceName     = "v2/rpc/user"
	GetUserInfoByUserId = "UserInfo"
)

// 服务标示
func (u *userService) ServiceName() string {
	return UserServiceName
}

type GetUserInfoReq struct {
	Type    string
	GroupId int32
	UserId  int32
}

type UserInfo struct {
	Name   string
	Avatar string
}

type GetUserInfoResp map[string]UserInfo

func (u *userService) GetUserInfo(ctx context.Context, req *GetUserInfoReq, resp *GetUserInfoResp) error {
	// 参数拼接
	params := []interface{}{
		req.Type,
		req.GroupId,
		req.UserId,
	}

	// 参数构造
	reqArgs := &rpc.ReqArgs{
		Method: GetUserInfoByUserId,
		Params: params,
	}

	// 请求
	err := rpc.Call(ctx, u, reqArgs, &resp)
	if err != nil {
		return err
	}

	return nil
}
