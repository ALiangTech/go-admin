package routes

import (
	"aliangtect/go-admin/routers/types"
	"aliangtect/go-admin/routers/v1/middlewares"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// 获取用户权限码

func RetrievePermission(router *gin.RouterGroup) {
	router.GET("/permission", func(ctx *gin.Context) {
		fmt.Println("permission")
		user, err := casbin.CasbinJsGetPermissionForUser(middlewares.Enforcer, "test")
		if err != nil {
			fmt.Println(err)
			return
		}
		ctx.JSON(200, types.ApiResponse{
			Data:  user,
			Error: nil,
		})
	})
}
