package routes

import "github.com/gin-gonic/gin"

// 获取用户权限码

func RetrievePermission(router *gin.RouterGroup) {
	router.GET("/permission", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"error": nil,
			"data": gin.H{
				"read": [1]int{0},
			},
		})
	})
}
