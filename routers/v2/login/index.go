package login

import (
	"aliangtect/go-admin/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Users struct {
	pswmatch bool
}

// request body JSON

type Form struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// 处理用户登录
var users Users

// db.DB.Raw("SELECT id, name FROM users").Scan(&users)
// fmt.Println(users)
func HanderLogin(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		form := Form{}
		if err := ctx.ShouldBind(&form); err == nil {
			db.DB.Raw("SELECT (pwd = crypt(?, pwd)) AS pswmatch FROM users where name = ?;", form.Pwd, form.Name).Scan(&users)
			fmt.Println(users.pswmatch)
		}
		ctx.JSON(200, gin.H{
			"message": "pong3",
		})
	})
}
