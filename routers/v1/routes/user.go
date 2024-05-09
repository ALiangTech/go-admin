package routes

import (
	"aliangtect/go-admin/db"
	"aliangtect/go-admin/routers/types"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 查询用户信息结构
type user struct {
	Name string `json:"name"`
}

// 注册用户信息结构
type register_user struct {
	Name   string `json:"name" binding:"required"`
	Pwd    string `json:"pwd" binding:"required"`
	RoleId int    `json:"roleId" binding:"required"`
}

func RetrieveUser(router *gin.RouterGroup) {
	// 获取当前登录用户信息
	router.GET("/profile", func(ctx *gin.Context) {
		uuid, exists := ctx.Get("userUuid")
		if exists {
			var user_info user
			// uuid 存在则去数据库查询
			res := db.DB.Raw("select * from users where uuid = ?", uuid).Scan(&user_info)
			if res.RowsAffected == 1 { // 说明查到了这个uuid 对应的用户数据
				fmt.Printf("%+v", res.Statement)
				ctx.JSON(http.StatusOK, types.ApiResponse{
					Data:  user_info,
					Error: nil,
				})
			} else {
				ctx.JSON(http.StatusOK, types.ApiResponse{
					Data: nil,
					Error: &types.ApiError{
						Status:     0,
						StatusText: "用户不存在",
					},
				})
			}
		} else {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Status:     0,
					StatusText: "用户不存在",
				},
			})
		}
	})
	// 注册用户
	router.POST("/user", func(ctx *gin.Context) {
		//获取表单
		// todo 判断账号是否存在
		// 插入数据库
		register_user_info := register_user{}
		err := ctx.ShouldBind(&register_user_info)
		if err == nil {
			res := db.DB.Exec("INSERT into users (name,pwd,roldid,createdOn) values (?, crypt(?, gen_salt('bf')),?,now())",
				register_user_info.Name, register_user_info.Pwd, register_user_info.RoleId)
			if res.RowsAffected == 1 {
				// 说明插入成功
				ctx.JSON(http.StatusOK, types.ApiResponse{
					Data: gin.H{
						"message": "注册成功",
					},
					Error: nil,
				})
			}
		}
		fmt.Println(err)

	})
	// 删除用户 软删除

	// 获取用户列表 可分页 可查询
	type userItem struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		Createdon time.Time `json:"createdOn"`
	}
	type totalItem struct {
		Pagecount int `json:"pageCount"`
	}
	router.GET("/users", func(ctx *gin.Context) {
		// 获取分页数据
		// size 数据
		// page 页码
		// query 查询
		// 先判断分页参数
		limit, _ := strconv.Atoi(ctx.Query("pageSize"))
		page, _ := strconv.Atoi(ctx.Query("page"))
		offset := limit * (page - 1)
		var data []userItem
		var total []totalItem
		db.DB.Raw("select count(*) as pagecount from users").Scan(&total) // 查询total
		db.DB.Raw("select * from users order by id limit ? offset ?", limit, offset).Scan(&data)
		ctx.JSON(http.StatusOK, types.ApiResponse{
			Data: types.Pagation[userItem]{
				Record:    data,
				PageCount: total[0].Pagecount,
			},
		})
	})
}
