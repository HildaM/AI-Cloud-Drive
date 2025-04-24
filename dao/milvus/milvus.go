package milvus

//go:generate $GOPATH/bin/mockgen --source=milvus.go --destination=milvus_mock.go --package=milvus

import "context"

type Dao interface {
	SaveChunks(ctx context.Context, chunks []Chunk) error
	Search(ctx context.Context, kbID string, vector []float32, topK int) ([]Chunk, error)
	DeleteChunks(ctx context.Context, docIDs []string) error
}
