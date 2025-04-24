package storage

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hildam/AI-Cloud-Drive/conf"
)

func init() {
	register(conf.LocalStorageType, &localImpl{})
}

type localImpl struct {
	baseDir string // 本地存储根目录（如 ./storage_data）
}

// New  初始化
func (l *localImpl) New(ctx context.Context, cfg *conf.StorageConfig) (Driver, error) {
	localCfg := cfg.Local
	// 确保存储目录存在
	if err := os.MkdirAll(localCfg.BaseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create local storage dir: %v", err)
	}
	return &localImpl{baseDir: localCfg.BaseDir}, nil
}

// Upload 上传文件
func (l *localImpl) Upload(ctx context.Context, data []byte, key string, contentType string) error {
	fullPath := filepath.Join(l.baseDir, key)
	// 确保父目录存在
	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return fmt.Errorf("failed to create parent dir: %v", err)
	}
	return os.WriteFile(fullPath, data, 0644)
}

// Download 下载文件
func (l *localImpl) Download(ctx context.Context, key string) ([]byte, error) {
	fullPath := filepath.Join(l.baseDir, key)
	return os.ReadFile(fullPath)
}

// Delete 删除文件
func (l *localImpl) Delete(ctx context.Context, key string) error {
	fullPath := filepath.Join(l.baseDir, key)
	return os.Remove(fullPath)
}

// GetURL 获取访问URL
func (l *localImpl) GetURL(ctx context.Context, key string) (string, error) {
	return filepath.Join(l.baseDir, key), nil
}
