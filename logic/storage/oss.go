package storage

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/hildam/AI-Cloud-Drive/conf"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func init() {
	register(conf.OssStorageType, &ossImpl{})
}

// ossImpl 阿里云OSS存储驱动结构体
type ossImpl struct {
	expireTime int64       // 过期时间（秒）
	bucket     *oss.Bucket // OSS Bucket实例
}

// New  初始化
func (o *ossImpl) New(ctx context.Context, cfg *conf.StorageConfig) (Driver, error) {
	ossCfg := cfg.OSS
	// 创建OSS客户端
	client, err := oss.New(ossCfg.Endpoint, ossCfg.AccessKeyID, ossCfg.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("failed to create OSS client: %v", err)
	}

	// 获取Bucket实例
	bucket, err := client.Bucket(ossCfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to get OSS bucket: %v", err)
	}
	return &ossImpl{bucket: bucket, expireTime: cfg.UrlExpireTime}, nil
}

// Upload 上传文件
func (o *ossImpl) Upload(ctx context.Context, data []byte, key string, contentType string) error {
	return o.bucket.PutObject(key, bytes.NewReader(data))
}

// Download 下载文件
func (o *ossImpl) Download(ctx context.Context, key string) ([]byte, error) {
	reader, err := o.bucket.GetObject(key)
	if err != nil {
		return nil, fmt.Errorf("failed to download from OSS: %v", err)
	}
	defer reader.Close()

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(reader); err != nil {
		return nil, fmt.Errorf("failed to read OSS data: %v", err)
	}
	return buf.Bytes(), nil
}

// Delete 删除文件
func (o *ossImpl) Delete(ctx context.Context, key string) error {
	return o.bucket.DeleteObject(key)
}

// GetURL 获取访问URL
func (o *ossImpl) GetURL(ctx context.Context, key string) (string, error) {
	expired := time.Now().Add(time.Duration(o.expireTime) * time.Second)
	return o.bucket.SignURL(key, oss.HTTPGet, int64(expired.Unix()))
}
