package routers

import (
	"aliangtect/go-admin/routers/v2/login"

	"github.com/gin-gonic/gin"
)

// 路由组v1 需要登录才能访问

var V1 *gin.RouterGroup

// 路由组v2 无需登录就能访问

var V2 *gin.RouterGroup

func BootGin() {
	router := gin.Default()
	V1 = router.Group("v1")
	{
		V1.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	V2 = router.Group("v2")
	{
		login.HanderLogin(V2)
	}
	router.Run(":9000")
}
