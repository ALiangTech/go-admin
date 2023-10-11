package middlewares

import "github.com/gin-gonic/gin"

// RBAC模型

const rbac_models = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`

func init() {

}

// 通过casbin 用来鉴定用户权限

// 通过jwt 获取角色

func validateUserPermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
