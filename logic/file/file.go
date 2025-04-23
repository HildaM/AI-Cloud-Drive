package file

//go:generate $GOPATH/bin/mockgen --source=file.go --destination=file_mock.go --package=file

import (
	"mime/multipart"

	"github.com/hildam/AI-Cloud-Drive/dao/file"
)

type Logic interface {
	// 上传文件
	UploadFile(userID uint, fileHeader *multipart.FileHeader, file multipart.File, parentID string) (string, error)
	// 获取文件URL
	GetFileURL(key string) (string, error)
	// 分页列表
	PageList(userID uint, parentID *string, page int, pageSize int, sort string) (int64, []file.File, error)
	// 下载文件
	DownloadFile(fileID string) (*file.File, []byte, error)
	// 删除文件或文件夹
	DeleteFileOrFolder(userID uint, fileID string) error
	// 创建文件夹
	CreateFolder(userID uint, name string, parentID *string) error
	// 批量移动文件
	BatchMoveFiles(userID uint, fileIDs []string, targetParentID string) error
	// 搜索列表
	SearchList(userID uint, key string, page int, size int, sort string) (int64, []file.File, error)
	// 重命名
	Rename(userID uint, fileID string, newName string) error
	// 获取文件路径
	GetFilePath(fileID string) (string, error)
	// 获取文件ID路径
	GetFileIDPath(fileID string) (string, error)
	// 获取文件
	GetFileByID(fileID string) (*file.File, error)
	// 初始化知识库目录
	InitKnowledgeDir(userID uint) (string, error)
}
