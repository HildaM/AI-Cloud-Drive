package storage

//go:generate $GOPATH/bin/mockgen --source=storage.go --destination=storage_mock.go --package=storage

import (
	"context"
	"fmt"

	"github.com/hildam/AI-Cloud-Drive/conf"
)

// API 定义存储驱动接口
type Driver interface {
	New(ctx context.Context, cfg *conf.StorageConfig) (Driver, error)              // 初始化
	Upload(ctx context.Context, data []byte, key string, contentType string) error // 上传文件
	Download(ctx context.Context, key string) ([]byte, error)                      // 下载文件
	Delete(ctx context.Context, key string) error                                  // 删除文件
	GetURL(ctx context.Context, key string) (string, error)                        // 获取访问URL
}

var driverMap = make(map[string]Driver)

// register 注册服务
func register(driverType string, driver Driver) {
	driverMap[driverType] = driver
}

// NewDriver 根据配置初始化存储驱动
func NewDriver(ctx context.Context, cfg *conf.StorageConfig) (Driver, error) {
	// 配置检查
	cfg.CheckCfg()

	// 获取实例
	if driver, ok := driverMap[cfg.Type]; ok {
		return driver.New(ctx, cfg)
	}
	return nil, fmt.Errorf("unsupported storage type: %s", cfg.Type)
}
