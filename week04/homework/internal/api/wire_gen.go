// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package api

import (
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/internal/biz"
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/internal/service"
)

// Injectors from wire.go:

func InitializeAdduser(username string) (int64, error) {
	user, err := biz.AddUser(username)
	if err != nil {
		return 0, err
	}
	int64_2, err := service.AddUser(user)
	if err != nil {
		return 0, err
	}
	return int64_2, nil
}
