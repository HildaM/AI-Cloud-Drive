package user

//go:generate $GOPATH/bin/mockgen --source=user.go --destination=user_mock.go --package=user

import (
	"github.com/hildam/AI-Cloud-Drive/dao/user"
)

type Logic interface {
	// 注册请求
	Register(user *user.User) error
	// 登录请求
	Login(req *LoginReq) (*LoginRsp, error)
}
