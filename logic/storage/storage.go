package storage

import (
	"fmt"

	"github.com/hildam/AI-Cloud-Drive/conf"
)

// Driver 定义存储驱动接口
type Driver interface {
	Upload(data []byte, key string, contentType string) error // 上传文件
	Download(key string) ([]byte, error)                      // 下载文件
	Delete(key string) error                                  // 删除文件
	GetURL(key string) (string, error)                        // 获取访问URL
}

// NewDriver 根据配置初始化存储驱动
// TODO 应该参考 executor 改为 “工厂模式“ 去实现
func NewDriver(cfg conf.StorageConfig) (Driver, error) {
	switch cfg.Type {
	case "local":
		return NewLocalStorage(cfg.Local.BaseDir)
	case "oss":
		return NewOSSStorage(cfg.OSS)
	case "minio":
		return NewMinioStorage(cfg.Minio)
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", cfg.Type)
	}
}
