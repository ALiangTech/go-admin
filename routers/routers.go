package routers

import (
	v1 "aliangtect/go-admin/routers/v1"
	v2 "aliangtect/go-admin/routers/v2"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BootGin() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "content-type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// v1 路由模块 需要登录才可以使用的接口
	v1.RegisterV1(router)
	// v2 路由模块 无需登录就可以使用的接口
	v2.RegisterV2(router)
	// 启动
	router.Run(":9001")
}
