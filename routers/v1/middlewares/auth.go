package middlewares

import (
	"aliangtect/go-admin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// jwt 登录鉴权

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从header里面获取jwt
		// 解析 jwt 并设置一个变量下去方便后续的接口从jwt中读取数据

		jwt := ctx.GetHeader("Authorization")

		if len(jwt) == 0 {
			// 说明没有携带jwt 这个时候返回认证失败
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "登录令牌认证失败",
				"data":    "",
			})
		} else {
			payload, err := utils.ParseJwt(jwt)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status":  http.StatusUnauthorized,
					"message": "登录令牌认证失败",
					"data":    "",
				})
			}
			ctx.Set("userUuid", payload.Uuid)
			ctx.Next()
		}

	}
}
