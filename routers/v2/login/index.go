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
		// 有错误 但是我依旧会返回基础结构给你 但是数据肯定是空
		// 没有错误 就正常返回数据
		if validateRequestContentType(gin.MIMEJSON, ctx) {
			// 请求体 content-type 类型类型正确
			// 请求体参数校验 字段个数 字段值类型
			err := parseBody(ctx, &form)
			if err := ctx.ShouldBindJSON(&form); err == nil {
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
			} else {
				fmt.Println(err)
				fmt.Println("err")
			}
		} else {
			ctx.JSON(200, gin.H{
				"code":    1,
				"message": fmt.Sprintf("请求体内容类型(content-type)错误,接口只能处理%v类型", gin.MIMEJSON),
				"data":    "",
			})
		}
	})
}

/**
* 校验请求进来的content-type 是否正确
* 参数：
*  - ctx: gin 框架上下文对象，用于获取请求信息
*  - target： 目标content-type 用于与请求的content-type 进行比较
* 返回值：
*  - 如果请求content-type 与 目标 相匹配, 则返回true; 否则返回false
 */
func validateRequestContentType(target string, ctx *gin.Context) bool {
	requestContentType := ctx.ContentType()
	return requestContentType == target
}

/**
* 解析请求体里面的body数据
* 参数：
*	  - ctx: Gin 上下文对象，用于处理 HTTP 请求。
*   - form: 指向表单结构体 (*Form) 的指针，用于存储解析后的数据。
* 返回值:
*   - error: 如果解析和绑定过程中出现错误，则返回一个描述错误的错误对象；否则返回 nil。
 */
func parseBody(ctx *gin.Context, form *Form) error {
	err := ctx.ShouldBindJSON(form)
	return err
}

func response(ctx *gin.Context) {

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
