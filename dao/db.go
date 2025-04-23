package dao

import (
	"fmt"

	"github.com/hildam/AI-Cloud-Drive/conf"
	"github.com/hildam/AI-Cloud-Drive/dao/file"
	"github.com/hildam/AI-Cloud-Drive/dao/knowledge"
	"github.com/hildam/AI-Cloud-Drive/dao/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbCli *gorm.DB

// InitDb 初始化 db
func InitDb() (err error) {
	// 构造 dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetCfg().Database.User,
		conf.GetCfg().Database.Password,
		conf.GetCfg().Database.Host,
		conf.GetCfg().Database.Port,
		conf.GetCfg().Database.Name,
	)

	// 创建实例
	dbCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// 自动迁移
	if err := dbCli.AutoMigrate(
		&user.User{},
		&file.File{},
		&knowledge.KnowledgeBase{},
		&knowledge.Document{},
	); err != nil {
		return err
	}
	return nil
}

// GetDb 获取DB实例
func GetDb() *gorm.DB {
	return dbCli
}
