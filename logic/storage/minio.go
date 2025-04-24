package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func init() {
	register(conf.MinioStorageType, &minioImpl{})
}

type minioImpl struct {
	client     *minio.Client
	bucket     string
	expireTime int64 // 单位秒
}

// New  初始化
func (m *minioImpl) New(ctx context.Context, cfg *conf.StorageConfig) (Driver, error) {
	minioCfg := cfg.Minio
	// 初始化 Minio 客户端
	client, err := minio.New(minioCfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioCfg.AccessKeyID, minioCfg.AccessKeySecret, ""),
		Secure: minioCfg.UseSSL,
		Region: minioCfg.Region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client: %v", err)
	}

	// 检查 bucket 是否存在
	exists, err := client.BucketExists(context.Background(), minioCfg.Bucket)
	if err != nil {
		return nil, fmt.Errorf("failed to check bucket existence: %v", err)
	}

	// 如果 bucket 不存在，创建它
	if !exists {
		err = client.MakeBucket(context.Background(), minioCfg.Bucket, minio.MakeBucketOptions{
			Region: minioCfg.Region,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create bucket: %v", err)
		}
	}
	return &minioImpl{
		client:     client,
		bucket:     minioCfg.Bucket,
		expireTime: cfg.UrlExpireTime,
	}, nil
}

// Upload 上传文件
func (m *minioImpl) Upload(ctx context.Context, data []byte, key string, contentType string) error {
	reader := bytes.NewReader(data)
	_, err := m.client.PutObject(
		context.Background(),
		m.bucket,
		key,
		reader,
		int64(len(data)),
		minio.PutObjectOptions{
			ContentType: contentType, // 例如 "application/pdf"
		},
	)
	if err != nil {
		return fmt.Errorf("failed to upload file: %v", err)
	}
	return nil
}

// Download 下载文件
func (m *minioImpl) Download(ctx context.Context, key string) ([]byte, error) {
	obj, err := m.client.GetObject(ctx, m.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %v", err)
	}
	defer obj.Close()

	data, err := io.ReadAll(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to read object data: %v", err)
	}
	return data, nil
}

// Delete 删除文件
func (m *minioImpl) Delete(ctx context.Context, key string) error {
	err := m.client.RemoveObject(ctx, m.bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}
	return nil
}

// GetURL 获取访问URL
func (m *minioImpl) GetURL(ctx context.Context, key string) (string, error) {
	// 设置响应头，强制浏览器下载文件
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment")

	// 生成预签名URL，有效期1小时
	expiry := time.Duration(m.expireTime) * time.Second
	presignedURL, err := m.client.PresignedGetObject(
		ctx,
		m.bucket,
		key,
		expiry,
		reqParams, // 关键：传递自定义参数
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}
	return presignedURL.String(), nil
}

// CreateDirectory 创建目录（通过上传空对象实现）
func (m *minioImpl) CreateDirectory(ctx context.Context, dirPath string) error {
	// 确保路径以 / 结尾
	if !strings.HasSuffix(dirPath, "/") {
		dirPath = dirPath + "/"
	}

	// 上传一个空对象来表示目录
	_, err := m.client.PutObject(ctx, m.bucket, dirPath, bytes.NewReader([]byte{}), 0, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	return nil
}
