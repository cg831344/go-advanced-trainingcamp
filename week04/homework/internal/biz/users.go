package biz

import (
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/internal/service"
	"github.com/pkg/errors"
)

// 检测用户名是否符合规范
func checkUserName(username string) error {
	u := []rune(username)
	if len(u) < 2 {
		return errors.New("姓名长度应大于2")
	}
	return nil
}

func AddUser(username string) (*service.User, error) {
	if err := checkUserName(username); err != nil {
		return nil, err
	}

	return &service.User{UserName: username}, nil

}
