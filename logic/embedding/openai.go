package embedding

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/hildam/AI-Cloud-Drive/conf"
)

func init() {
	register(conf.OllamaEmbeddingType, &openaiImpl{})
}

type openaiImpl struct {
	embedder *openai.Embedder
}

// New 创建实例
func (o *openaiImpl) New(ctx context.Context, cfg *conf.EmbeddingConfig) (Logic, error) {
	embedder, err := openai.NewEmbedder(ctx, &openai.EmbeddingConfig{
		APIKey:     cfg.Remote.APIKey,
		Model:      cfg.Remote.Model,
		BaseURL:    cfg.Remote.BaseURL,
		Timeout:    time.Duration(cfg.Remote.Timeout) * time.Second,
		Dimensions: &cfg.Remote.Dimension,
	})
	if err != nil {
		return nil, err
	}
	return &openaiImpl{embedder: embedder}, nil
}

// EmbedStrings 将文本转换为向量表示
func (o *openaiImpl) EmbedStrings(ctx context.Context, texts []string) ([][]float64, error) {
	return o.embedder.EmbedStrings(ctx, texts)
}
