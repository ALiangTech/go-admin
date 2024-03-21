package routesV2

import (
	"aliangtect/go-admin/db"
	cerrors "aliangtect/go-admin/error"
	"aliangtect/go-admin/routers/types"
	"aliangtect/go-admin/utils"
	"errors"
	"net/http"

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
		inputForm := Form{}
		err := ctx.ShouldBindJSON(&inputForm)
		if err != nil {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Code:    cerrors.ErrUser,
					Message: err.Error(),
				},
			})
			return
		}
		if inputForm.Name == "" || inputForm.Pwd == "" {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Code:    cerrors.ErrEmptyCredentials,
					Message: cerrors.StatusText(cerrors.ErrEmptyCredentials),
				},
			})
			return
		}
		// sql 认证用户登录
		err = AuthenticateUser(inputForm, &users)
		if err != nil {
			errorMsg := types.ApiError{
				Code:    cerrors.ErrUser,
				Message: err.Error(),
			}
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data:  nil,
				Error: &errorMsg,
			})
			return
		}
		jwt, err := utils.GenerateJwt(users.Uuid)
		if err != nil {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Code:    cerrors.ErrGenJwt,
					Message: cerrors.StatusText(cerrors.ErrGenJwt),
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, types.ApiResponse{
			Data: gin.H{
				"jwt": jwt,
			},
			Error: nil,
		})
	})
}

// 从结构体获取参数
// 校验从结构体获取的参数
// sql认证
// 返回结果给前端

func AuthenticateUser(inputForm Form, userRecord *Users) error {
	// 准备预处理语句
	stmt := db.DB.Raw("SELECT (pwd = crypt(?, pwd)) AS pswmatch, uuid FROM users WHERE name = ?;", inputForm.Pwd, inputForm.Name).Scan(userRecord)

	// 执行查询
	if err := stmt.Error; err != nil {
		return err
	}

	// 检查密码是否匹配
	if !userRecord.Pswmatch {
		return errors.New(cerrors.StatusText(cerrors.ErrUser))
	}

	return nil
}
