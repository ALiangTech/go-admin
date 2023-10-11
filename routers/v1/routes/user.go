package user

import "github.com/gin-gonic/gin"

// 查询用户信息

func RetrieveUser(router *gin.RouterGroup) {
	router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
