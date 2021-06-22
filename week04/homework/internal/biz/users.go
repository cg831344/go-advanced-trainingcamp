package biz

import (
	"github.com/pkg/errors"
)

type User struct {
	UserName string
}

// 检测用户名是否符合规范
func checkUserName(username string) error {
	u := []rune(username)
	if len(u) < 2 {
		return errors.New("姓名长度应大于2")
	}
	return nil
}

func AddUser(username string) (*User, error) {
	if err := checkUserName(username); err != nil {
		return nil, err
	}

	return &User{UserName: username}, nil

}
