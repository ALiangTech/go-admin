package routesV2

import (
	cerrors "aliangtect/go-admin/error"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Pswmatch bool
	Uuid     string
}

// request body JSON

type Form struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// 处理用户登录
var users Users

func RetrieveLogin(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		// 接口正常流程 从结构体获取参数 => 参数校验(不为空校验) => sql认证 => 返回结果给前端
		users := Form{}
		ctx.ShouldBindJSON(&users)
		if users.Name == "" || users.Pwd == "" {
			ctx.JSON(cerrors.ErrEmptyCredentials, gin.H{
				"data": nil,
				"error": gin.H{
					"code":    1,
					"message": cerrors.StatusText(cerrors.ErrEmptyCredentials),
				},
			})
		}
	})
}

// 从结构体获取参数
// 校验从结构体获取的参数
// sql认证
// 返回结果给前端
