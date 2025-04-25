package embedding

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/ollama/ollama/api"
)

func init() {
	register(conf.OllamaEmbeddingType, &ollamaImpl{})
}

type ollamaImpl struct {
	cli   *api.Client
	model string
}

// New 创建一个新的ollama实例
func (o *ollamaImpl) New(ctx context.Context, cfg *conf.EmbeddingConfig) (Logic, error) {
	// 检查配置
	config := cfg.Ollama
	cfg.CheckCfg()

	// 构造 cli
	httpCli := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second,
	}
	// 构造 url
	baseURL, err := url.Parse(config.Url)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	// 返回实例
	return &ollamaImpl{
		cli:   api.NewClient(baseURL, httpCli),
		model: config.Model,
	}, nil
}

// EmbedStrings 使用ollama的embed模型生成文本的向量表示
func (o *ollamaImpl) EmbedStrings(ctx context.Context, texts []string) ([][]float64, error) {
	// 构造请求
	resp, err := o.cli.Embed(ctx, &api.EmbedRequest{
		Model: o.model,
		Input: texts,
	})
	if err != nil {
		return nil, err
	}

	// 解析数据
	embeddings := make([][]float64, len(resp.Embeddings))
	for i, d := range resp.Embeddings {
		res := make([]float64, len(d))
		for j, emb := range d {
			res[j] = float64(emb)
		}
		embeddings[i] = res
	}
	return embeddings, nil
}
