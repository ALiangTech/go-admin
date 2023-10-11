package routers

import (
	v1 "aliangtect/go-admin/routers/v1"
	"aliangtect/go-admin/routers/v2/login"

	"github.com/gin-gonic/gin"
)

// 路由组v2 无需登录就能访问

var V2 *gin.RouterGroup

func BootGin() {
	router := gin.Default()
	V2 = router.Group("v2")
	// v1 路由模块
	{
		v1.RegisterV1(router)
	}
	// v2 路由模块
	{
		login.HanderLogin(V2)
	}
	router.Run(":9001")
}
