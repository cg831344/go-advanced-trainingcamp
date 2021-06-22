package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/cg831344/go-advanced-trainingcamp/week04/homework/internal/api"
)

func InitAddUser(router *gin.RouterGroup) {
	router.POST("addUser", api.Adduser)
}
