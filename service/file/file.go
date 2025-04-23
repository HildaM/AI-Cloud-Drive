package file

import (
	"github.com/hildam/AI-Cloud-Drive/logic/file"
	"github.com/labstack/echo/v4"
)

type fileSerivce struct {
	fileLogic file.Logic
}

func NewFileService() *fileSerivce {
	return &fileSerivce{
		fileLogic: file.NewFileLogic(),
	}
}

// Register 注册接口
func (u *fileSerivce) Upload(e echo.Context) error {
	return nil
}

// Login 登录接口

func (u *fileSerivce) PageList(e echo.Context) error {
	return nil
}

func (u *fileSerivce) Download(e echo.Context) error {
	return nil
}

func (u *fileSerivce) Delete(e echo.Context) error {
	return nil
}
func (u *fileSerivce) CreateFolder(e echo.Context) error {
	return nil
}
func (u *fileSerivce) BatchMove(e echo.Context) error {
	return nil
}
func (u *fileSerivce) Search(e echo.Context) error {
	return nil
}
func (u *fileSerivce) Rename(e echo.Context) error {
	return nil
}
func (u *fileSerivce) GetPath(e echo.Context) error {
	return nil
}
func (u *fileSerivce) GetIDPath(e echo.Context) error {
	return nil
}
