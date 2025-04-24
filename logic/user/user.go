package user

//go:generate $GOPATH/bin/mockgen --source=user.go --destination=user_mock.go --package=user

import (
	"context"
	"github.com/hildam/AI-Cloud-Drive/dao/user"
)

type Logic interface {
	// 注册请求
	Register(ctx context.Context, user *user.User) error
	// 登录请求
	Login(ctx context.Context, req *LoginReq) (*LoginRsp, error)
}
