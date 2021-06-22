package service

import (
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/database"
	"github.com/pkg/errors"
)

type User struct {
	UserName string
}

func AddUser(user *User) (int64, error) {
	res, err := database.DB.Exec("insert INTO users(name,age) values(?)", user.UserName)
	if err != nil {
		return 0, errors.Wrap(err, "新增用户失败")
	}
	id, _ := res.LastInsertId()
	return id, nil
}
