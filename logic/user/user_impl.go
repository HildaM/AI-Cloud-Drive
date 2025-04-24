package user

import (
	"context"
	"errors"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/hildam/AI-Cloud-Drive/dao"
	"github.com/hildam/AI-Cloud-Drive/dao/user"
	"github.com/hildam/AI-Cloud-Drive/service/middleware"
	"golang.org/x/crypto/bcrypt"
)

type userLogic struct {
	userDao user.Dao
}

func NewUserLogic(ctx context.Context) Logic {
	return &userLogic{
		userDao: user.NewUserDao(dao.GetDb()),
	}
}

func (s *userLogic) Register(ctx context.Context, u *user.User) error {
	// 检查
	userExists, err := s.userDao.CheckIfExist("username", u.Username)
	if err != nil {
		return err
	}
	if userExists {
		return errors.New("用户名已注册")
	}

	// 手机是否存在
	phoneExists, err := s.userDao.CheckIfExist("phone", u.Phone)
	if err != nil {
		return err
	}
	if phoneExists {
		return errors.New("手机号已注册")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}

	// 创建用户
	if err = s.userDao.CreateUser(&user.User{
		Username: u.Username,
		Phone:    u.Phone,
		Password: string(hashedPassword),
		Email:    u.Email,
	}); err != nil {
		return err
	}
	return nil
}

func (s *userLogic) Login(ctx context.Context, req *LoginReq) (*LoginRsp, error) {
	user, err := s.userDao.GetUserByName(req.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	accessToken, err := middleware.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("系统错误")
	}

	return &LoginRsp{
		AccessToken: accessToken,
		ExpiresIn:   conf.GetCfg().JWT.ExpirationHours * 3600,
		TokenType:   "Bearer",
	}, nil
}
