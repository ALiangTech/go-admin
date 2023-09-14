package login

import (
	"aliangtect/go-admin/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Users struct {
	ID   int
	Name string
}

// 处理用户登录
var users Users

func HanderLogin(router *gin.RouterGroup) {
	router.GET("/login", func(ctx *gin.Context) {
		db.DB.Raw("SELECT * FROM users").Scan(&users)
		fmt.Print(users)
		ctx.JSON(200, gin.H{
			"message": "pong3",
		})
	})
}
