package user

//go:generate $GOPATH/bin/mockgen --source=user.go --destination=user_mock.go --package=user

type Dao interface {
	// 检查是否存在
	CheckIfExist(field string, value interface{}) (bool, error)
	// 创建用户
	CreateUser(user *User) error
	// 获取用户
	GetUserByName(name string) (user *User, err error)
}
