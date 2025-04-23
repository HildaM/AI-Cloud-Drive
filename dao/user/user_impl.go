package user

import (
	"gorm.io/gorm"
)

type userDao struct {
	db *gorm.DB // db实例
}

func NewUserDao(db *gorm.DB) Dao {
	return &userDao{db: db}
}

// CheckIfExist 检查是否存在
func (ud *userDao) CheckIfExist(field string, value interface{}) (bool, error) {
	var count int64
	if err := ud.db.Model(&User{}).Where(field+" = ?", value).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// CreateUser 创建用户
func (ud *userDao) CreateUser(user *User) error {
	result := ud.db.Create(user)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetUserByName 获取用户
func (ud *userDao) GetUserByName(name string) (*User, error) {
	var user User
	result := ud.db.Model(&User{}).Where("username = ?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
