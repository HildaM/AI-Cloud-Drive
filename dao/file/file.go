package file

//go:generate $GOPATH/bin/mockgen --source=file.go --destination=file_mock.go --package=file

// Dao 定义了文件操作的接口
type Dao interface {
	// 创建文件
	CreateFile(file *File) error
	// 获取父目录文件列表
	GetFilesByParentID(userID uint, parentID *string) ([]File, error)
	// 获取文件元信息
	GetFileMetaByFileID(id string) (*File, error)
	// 删除文件
	DeleteFile(id string) error
	// 分页列表
	ListFiles(userID uint, parentID *string, page int, pageSize int, sort string) ([]File, error)
	// 获取父目录文件数量
	CountFilesByParentID(parentID *string, userID uint) (int64, error)
	// 更新文件
	UpdateFile(file *File) error
	// 获取关键词文件数量
	CountFilesByKeyword(key string, userID uint) (int64, error)
	// 获取关键词文件列表
	GetFilesByKeyword(userID uint, key string, page int, pageSize int, sort string) ([]File, error)
	// 获取知识库目录
	GetDocumentDir(userID uint) (*File, error)
}
