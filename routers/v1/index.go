package v1

import (
	"aliangtect/go-admin/routers/v1/middlewares"
	"aliangtect/go-admin/routers/v1/routes"
	"github.com/gin-gonic/gin"
)

// V1 路由组v1 需要登录才能访问
var V1 *gin.RouterGroup

func RegisterV1(router *gin.Engine) {
	V1 = router.Group("v1")
	V1.Use(middlewares.AuthMiddleware())
	V1.Use(middlewares.ValidateUserPermissions())
	routes.RetrieveUser(V1)
	routes.RetrieveRole(V1)
	routes.RetrievePermission(V1)
}
