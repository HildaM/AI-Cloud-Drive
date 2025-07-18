package knowledge

import (
	"time"
)

// KnowledgeBase 知识库
type KnowledgeBase struct {
	ID          string    `gorm:"primaryKey;type:char(36)"` // UUID
	Name        string    `gorm:"not null"`                 // 知识库名称
	Description string    // 知识库描述
	UserID      uint      `gorm:"index"` // 创建者ID
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// Document 知识库文档
type Document struct {
	ID              string    `gorm:"primaryKey;type:char(36)"` // UUID
	UserID          uint      `gorm:"index"`                    // 所属的用户
	KnowledgeBaseID string    `gorm:"index"`                    // 所属知识库ID
	FileID          string    `gorm:"index"`                    // 关联的文件ID
	Title           string    // 文档标题
	DocType         string    // 文档类型(pdf/txt/md)
	Status          int       // 处理状态(0:待处理,1:处理中,2:已完成,3:失败)
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
