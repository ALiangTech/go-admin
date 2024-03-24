package routes

import (
	"aliangtect/go-admin/routers/types"
	"aliangtect/go-admin/routers/v1/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取用户权限码

func RetrievePermission(router *gin.RouterGroup) {
	router.GET("/permission", func(ctx *gin.Context) {
		fmt.Println("permission")
		permissions, err := middlewares.Enforcer.GetImplicitPermissionsForUser("test")
		if err != nil {
			fmt.Println(err)
		}
		var permissionsCode []string
		for _, permission := range permissions {
			permissionsCode = append(permissionsCode, permission[0])
		}
		ctx.JSON(200, types.ApiResponse{
			Data:  permissionsCode,
			Error: nil,
		})
	})
}
