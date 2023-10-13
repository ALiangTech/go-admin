package routes

import "github.com/gin-gonic/gin"

func RetrieveRole(router *gin.RouterGroup) {
	// 创建角色
	router.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

// 根据权限表 构建权限树
