package login

import (
	"aliangtect/go-admin/db"
	"aliangtect/go-admin/utils"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Pswmatch bool
}

// request body JSON

type Form struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

// 处理用户登录
var users Users

func HanderLogin(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		form := Form{}
		if err := ctx.ShouldBind(&form); err == nil {
			message := validateInputNotEmpty(form)
			if len(message) != 0 { // 说明存在错误信息
				ctx.JSON(200, gin.H{
					"code":    1,
					"message": message,
					"data":    "",
				})
			} else {
				validateInputPwd(form, &users)
				if users.Pswmatch { // 密码正确
					token, _ := utils.GenerateJwt()
					ctx.JSON(200, gin.H{
						"code":    200,
						"message": "",
						"data": gin.H{
							"toke": token,
						},
					})
				} else {
					ctx.JSON(200, gin.H{
						"code":    1,
						"message": createMessage(),
						"data":    "",
					})
				}

			}
		}
	})
}

// 随机返回 账号 密码 中文字符串 顺序
func randName() (string, string) {
	const size = 2
	messageSet := [size]string{"账号", "密码"}
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	firstIndex := random.Intn(size) // 0 1
	lastIndex := size - firstIndex - 1
	return messageSet[firstIndex], messageSet[lastIndex]
}

// 构建错误信息字符串
func createMessage() string {
	m1, m2 := randName()
	return fmt.Sprintf("%s或%s错误", m1, m2)
}

// 检测用户输入的数据是否为空
func validateInputNotEmpty(form Form) string {
	if len(form.Name) == 0 || len(form.Pwd) == 0 { // 说明用户输入的数据 有问题
		return createMessage()
	}
	return ""
}

// 执行sql 校验密码是否正确 并查找用户

func validateInputPwd(form Form, users *Users) {
	db.DB.Raw("SELECT (pwd = crypt(?, pwd)) AS pswmatch FROM users where name = ?;", form.Pwd, form.Name).Scan(users)
}
