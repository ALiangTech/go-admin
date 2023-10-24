package routes

import (
	"aliangtect/go-admin/db"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type Role struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Policy      []string `json:"policy"`
}

// 向数据库插入数据

func interRole(role Role) {
	db.DB.Exec("INSERT INTO roles (name,description,policy) values (?, ?, ?)", role.Name, role.Description, pq.Array(role.Policy))
}

func RetrieveRole(router *gin.RouterGroup) {
	// 创建角色
	router.POST("/role", func(ctx *gin.Context) {
		role := Role{}
		ctx.ShouldBindJSON(&role)
		interRole(role)
		// 获取角色name 权限policy
		ctx.JSON(200, gin.H{
			"message": "pong",
			"role":    role,
		})
	})
}

// 根据权限表 构建权限树
