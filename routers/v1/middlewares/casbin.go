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

const rbacModels = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "root"
`

var Enforcer *casbin.Enforcer

// HOST     = "localhost"
// PORT     = 9002
// USER     = "goadmin"
// PASSWORD = "1234"
// NAME     = "admin"
// p,account,/account/*, (get|post)
// p,account_add, /account_add, post
// p,account_del, /account_del, post
// p,role,/role/*, (get|post)
// p,role_add, /role_add, post
// p,role_del, /role_del, del
// g,admin,account
// g,admin,role_del

func init() {
	fmt.Println("casbin init")
	a, err := gormadapter.NewAdapter("postgres", db.DSN, true)
	if err != nil {
		panic(err)
	}
	m, _ := model.NewModelFromString(rbacModels)
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	rules := [][]string{
		{"account", "/account/*", "(get|post)"}, // 账号管理
		{"account_add", "/account_add", "post"}, // 账号添加
		{"account_del", "/account_del", "post"}, // 账号删除
		{"permission", "/permission", "get"},    // 获取权限
	}
	_, err = e.AddPolicies(rules)
	if err != nil {
		panic(err)
	}
	// e.AddGroupingPolicy()
	err = e.LoadPolicy()
	if err != nil {
		return
	}
	err = e.SavePolicy()
	Enforcer = e
	if err != nil {
		return
	}
}

// ValidateUserPermissions 通过casbin 用来鉴定用户权限
func ValidateUserPermissions() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		vaild, _ := Enforcer.Enforce("admin", "/permission", "get")
		fmt.Println(vaild, "vaild")
		ctx.Next()

	}
}
