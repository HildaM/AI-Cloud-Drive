package embedding

import (
	"context"
	"fmt"

	"github.com/hildam/AI-Cloud-Drive/conf"
)

type Logic interface {
	New(ctx context.Context, cfg *conf.EmbeddingConfig) (Logic, error)
	// EmbedStrings 将文本转换为向量表示
	EmbedStrings(ctx context.Context, texts []string) ([][]float64, error)
}

var embeddingMap = make(map[string]Logic)

func register(name string, logic Logic) {
	embeddingMap[name] = logic
}

// NewEmbedding 根据配置初始化嵌入模型
func NewEmbedding(ctx context.Context, cfg *conf.EmbeddingConfig) (Logic, error) {
	// 配置检查
	if err := cfg.CheckCfg(); err != nil {
		return nil, err
	}

	// 获取实例
	if embedding, ok := embeddingMap[cfg.Type]; ok {
		return embedding.New(ctx, cfg)
	}
	return nil, fmt.Errorf("unsupported embedding type: %s", cfg.Type)
}
