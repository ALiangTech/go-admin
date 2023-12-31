package login

import (
	"aliangtect/go-admin/db"
	"aliangtect/go-admin/utils"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

func HanderLogin(router *gin.RouterGroup) {
	router.POST("/login", func(ctx *gin.Context) {
		// form := Form{}
		// 1 请求内容类型是否正确
		// 2 请求内容数据能否正确解析
		// 3 请求内容是否正确
		// 4 创建jwt 给前端
		validContentType := &ValidContentType{}
		validBody := &ValidBody{}
		validBodyFieldValue := &ValidBodyFieldValue{}
		validBodyCorrectness := &ValidBodyCorrectness{}
		createJwt := &CreateJwt{}
		validContentType.SetNext(validBody)
		validBody.SetNext(validBodyFieldValue)
		validBodyFieldValue.SetNext(validBodyCorrectness)
		validBodyCorrectness.SetNext(createJwt)
		validContentType.Execute(ctx)
	})
}

// 通过责任链 清除嵌套式if else 结构

// 校验基础结构
type Valid interface {
	Execute(*gin.Context)
	SetNext(Valid)
}

// 校验请求体content-type 是否符合要求
type ValidContentType struct {
	next Valid
}

/**
* 校验请求进来的content-type 是否正确
* 参数：
*  - ctx: gin 框架上下文对象，用于获取请求信息
*  - target： 目标content-type 用于与请求的content-type 进行比较
* 返回值：
*  - 如果请求content-type 与 目标 相匹配, 则返回true; 否则返回false
 */
func validateRequestContentType(ctx *gin.Context, target string) bool {
	requestContentType := ctx.ContentType()
	return requestContentType == target
}
func (vc *ValidContentType) Execute(ctx *gin.Context) {
	isValidContentType := validateRequestContentType(ctx, gin.MIMEJSON)
	if isValidContentType {
		vc.next.Execute(ctx)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  10001,
			"message": "请求体内容类型错误",
			"data":    "",
		})
	}
}
func (vc *ValidContentType) SetNext(next Valid) {
	vc.next = next
}

////////////////////////////////////////////

// 校验请求体body 数据字段是否完整

type ValidBody struct {
	next Valid
}

func (vb *ValidBody) Execute(ctx *gin.Context) {
	form := Form{}
	err := ctx.ShouldBindBodyWith(&form, binding.JSON)
	if err == nil {
		vb.next.Execute(ctx)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  10002,
			"message": fmt.Sprintf("请求的数据：%v", err),
			"data":    "",
		})
	}
}

func (vb *ValidBody) SetNext(next Valid) {
	vb.next = next
}

// ** 校验请求body 数据值是否正确

type ValidBodyFieldValue struct {
	next Valid
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
func validateInputNotEmpty(form Form) (string, error) {
	if len(form.Name) == 0 || len(form.Pwd) == 0 { // 说明用户输入的数据 有问题
		return createMessage(), errors.New("数据值有问题")
	}
	return "", nil
}
func (vf *ValidBodyFieldValue) Execute(ctx *gin.Context) {
	form := Form{}
	ctx.ShouldBindBodyWith(&form, binding.JSON)
	fmt.Println(form, "form")
	msg, err := validateInputNotEmpty(form)
	if err == nil {
		vf.next.Execute(ctx)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  10003,
			"message": fmt.Sprintf("请求数据值：%v", msg),
			"data":    "",
		})
	}
}

func (vf *ValidBodyFieldValue) SetNext(next Valid) {
	vf.next = next
}

////////////////////////////////////////////

// 校验请求体数据 正确性

type ValidBodyCorrectness struct {
	next Valid
}

// 执行sql 校验密码是否正确 并查找用户
func validateInputPwd(form Form, users *Users) {
	db.DB.Raw("SELECT (pwd = crypt(?, pwd)) AS pswmatch, uuid FROM users where name = ?;", form.Pwd, form.Name).Scan(users)
}

func (vbc *ValidBodyCorrectness) Execute(ctx *gin.Context) {
	form := Form{}
	ctx.ShouldBindBodyWith(&form, binding.JSON)
	validateInputPwd(form, &users)
	fmt.Println(users, "users")
	if users.Pswmatch {
		vbc.next.Execute(ctx)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  10004,
			"message": createMessage(),
			"data":    "",
		})
	}
}

func (vbc *ValidBodyCorrectness) SetNext(next Valid) {
	vbc.next = next
}

////////////////////////////////////////////
// 创建jwt

type CreateJwt struct {
	next Valid
}

func (cj *CreateJwt) Execute(ctx *gin.Context) {
	jwt, err := utils.GenerateJwt(users.Uuid)
	if err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": "",
			"data": gin.H{
				"token": jwt,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  10005,
			"message": "创建jwt失败",
			"data":    "",
		})
	}
}
func (cj *CreateJwt) SetNext(next Valid) {
	cj.next = next
}
