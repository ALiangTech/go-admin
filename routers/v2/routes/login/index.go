package routesV2

import (
	"aliangtect/go-admin/db"
	cerrors "aliangtect/go-admin/error"
	"aliangtect/go-admin/routers/types"
	"aliangtect/go-admin/routers/v1/routes"
	"aliangtect/go-admin/utils"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Users struct {
	Pswmatch bool
	Uuid     string
	RoleId   int
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
					Status:     cerrors.ErrUser,
					StatusText: err.Error(),
				},
			})
			return
		}
		if inputForm.Name == "" || inputForm.Pwd == "" {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Status:     cerrors.ErrEmptyCredentials,
					StatusText: cerrors.StatusText(cerrors.ErrEmptyCredentials),
				},
			})
			return
		}
		// sql 认证用户登录
		err = AuthenticateUser(inputForm, &users)
		if err != nil {
			errorMsg := types.ApiError{
				Status:     cerrors.ErrUser,
				StatusText: err.Error(),
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
					Status:     cerrors.ErrGenJwt,
					StatusText: cerrors.StatusText(cerrors.ErrGenJwt),
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
// sql 返回用户信息 然后用户信息存放在redis中
// 返回结果给前端
var this = context.TODO()

func AuthenticateUser(inputForm Form, userRecord *Users) error {
	// 准备预处理语句
	stmt := db.DB.Raw("SELECT (pwd = crypt(?, pwd)) AS pswmatch, uuid, roleId FROM users WHERE name = ?;", inputForm.Pwd, inputForm.Name).Scan(userRecord)
	// 执行查询
	if err := stmt.Error; err != nil {
		return err
	}

	// 检查密码是否匹配
	if !userRecord.Pswmatch {
		return errors.New(cerrors.StatusText(cerrors.ErrUser))
	}
	// 通过角色id 查询用户角色
	role, err := routes.GetRoleWithID(userRecord.RoleId)
	if err != nil {
		panic(err)
	}
	userInfo := map[string]any{
		"roleId":   userRecord.RoleId,
		"pswmatch": userRecord.Pswmatch,
		"roleName": role.Name,
	}
	err = db.Redis.HSet(this, userRecord.Uuid, userInfo).Err() // 用户信息存放在redis
	if err != nil {
		panic(err)
	}
	return nil
}
