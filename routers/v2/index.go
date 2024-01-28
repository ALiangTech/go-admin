package v2

import (
	routesV2 "aliangtect/go-admin/routers/v2/routes/login"

	"github.com/gin-gonic/gin"
)

// 路由组v2 无需登录才能访问
var V2 *gin.RouterGroup

func RegisterV2(router *gin.Engine) {
	V2 = router.Group("v2")
	routesV2.RetrieveLogin(V2)
}
