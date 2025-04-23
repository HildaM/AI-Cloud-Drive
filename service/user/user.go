package user

import (
	"github.com/hildam/AI-Cloud-Drive/logic/user"
	"github.com/labstack/echo/v4"
)

type userSerivce struct {
	userLogic user.Logic
}

func NewUserService() *userSerivce {
	return &userSerivce{
		userLogic: user.NewUserLogic(),
	}
}

// Register 注册接口
func (u *userSerivce) Register(e echo.Context) error {
	return nil
}

// Login 登录接口

func (u *userSerivce) Login(e echo.Context) error {
	return nil
}
