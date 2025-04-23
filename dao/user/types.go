package user

import (
	"time"
)

// User 用户信息
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"uniqueIndex;size:50;not null"`
	Phone     string    `gorm:"uniqueIndex;size:20;not null"` // 新增手机号字段
	Email     string    `gorm:"uniqueIndex;size:100"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
