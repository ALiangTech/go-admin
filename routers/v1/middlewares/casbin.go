package middlewares

import (
	"aliangtect/go-admin/db"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

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

var Enforcer *casbin.Enforcer

// HOST     = "localhost"
// PORT     = 9002
// USER     = "goadmin"
// PASSWORD = "1234"
// NAME     = "admin"

func init() {
	a, err := gormadapter.NewAdapter("postgres", db.DSN, true)
	if err != nil {
		panic(err)
	}
	m, _ := model.NewModelFromString(rbac_models)
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	rules := [][]string{
		{"user", "/user", "get"},
	}
	_, err = e.AddPolicies(rules)
	if err != nil {
		panic(err)
	}
	// e.AddGroupingPolicy()
	e.LoadPolicy()
	Enforcer = e
}

// 通过casbin 用来鉴定用户权限

// 通过jwt 获取角色

func ValidateUserPermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vaild, _ := Enforcer.Enforce("alice", "data1", "read")
		fmt.Println(vaild, "vaild")
		ctx.Next()

	}
}
