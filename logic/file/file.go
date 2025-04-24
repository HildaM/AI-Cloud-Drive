package file

//go:generate $GOPATH/bin/mockgen --source=file.go --destination=file_mock.go --package=file

import (
	"context"
	"mime/multipart"

	"github.com/hildam/AI-Cloud-Drive/dao/file"
)

type Logic interface {
	// 上传文件
	UploadFile(ctx context.Context, userID uint, fileHeader *multipart.FileHeader, file multipart.File, parentID string) (string, error)
	// 获取文件URL
	GetFileURL(ctx context.Context, key string) (string, error)
	// 分页列表
	PageList(ctx context.Context, userID uint, parentID *string, page int, pageSize int, sort string) (int64, []file.File, error)
	// 下载文件
	DownloadFile(ctx context.Context, fileID string) (*file.File, []byte, error)
	// 删除文件或文件夹
	DeleteFileOrFolder(ctx context.Context, userID uint, fileID string) error
	// 创建文件夹
	CreateFolder(ctx context.Context, userID uint, name string, parentID *string) error
	// 批量移动文件
	BatchMoveFiles(ctx context.Context, userID uint, fileIDs []string, targetParentID string) error
	// 搜索列表
	SearchList(ctx context.Context, userID uint, key string, page int, size int, sort string) (int64, []file.File, error)
	// 重命名
	Rename(ctx context.Context, userID uint, fileID string, newName string) error
	// 获取文件路径
	GetFilePath(ctx context.Context, fileID string) (string, error)
	// 获取文件ID路径
	GetFileIDPath(ctx context.Context, fileID string) (string, error)
	// 获取文件
	GetFileByID(ctx context.Context, fileID string) (*file.File, error)
	// 初始化知识库目录
	InitKnowledgeDir(ctx context.Context, suserID uint) (string, error)
}
