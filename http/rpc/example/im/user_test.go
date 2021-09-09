package im

import (
	"context"
	"fmt"
	"testing"
)

func Test_userService_GetUserInfo(t *testing.T) {
	type args struct {
		ctx  context.Context
		req  *GetUserInfoReq
		resp *GetUserInfoResp
	}
	resp := new(GetUserInfoResp)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"测试IM User",
			args{
				ctx: context.Background(),
				req: &GetUserInfoReq{
					Type:    "group",
					GroupId: 3027668,
					UserId:  78295309,
				},
				resp: resp,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewUserService()
			if err := s.GetUserInfo(tt.args.ctx, tt.args.req, tt.args.resp); (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println(tt.args.resp)
		})
	}
}
