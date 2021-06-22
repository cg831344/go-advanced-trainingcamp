package api

import (
	"github.com/gin-gonic/gin"
)

func Adduser(c *gin.Context) {
	username := c.PostForm("username")
	// 使用wire自动生成依赖
	_, err := InitializeAdduser(username)
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "",
		})
		return
	}
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "创建成功",
	})
}
