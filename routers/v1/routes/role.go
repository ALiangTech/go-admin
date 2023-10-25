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

func interRole(role Role) error {
	res := db.DB.Exec("INSERT INTO roles (name,description,policy) values (?, ?, ?)", role.Name, role.Description, pq.Array(role.Policy))
	return res.Error
}

// 删除数据库数据

func delRole(id string) error {
	res := db.DB.Exec("DELETE FROM roles where id = ? ", id)
	return res.Error
}

func RetrieveRole(router *gin.RouterGroup) {
	// 创建角色
	router.POST("/role", func(ctx *gin.Context) {
		role := Role{}
		ctx.ShouldBindJSON(&role)
		err := interRole(role)
		if err != nil {
			// 获取角色name 权限policy
			ctx.JSON(200, gin.H{
				"message": "角色创建成功",
				"role":    role,
			})
		} else {
			// 获取角色name 权限policy
			ctx.JSON(200, gin.H{
				"message": "角色创建失败",
				"role":    role,
			})
		}
	})
	// 删除角色
	router.DELETE("/role/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := delRole(id)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "删除成功",
				"id":      id,
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": "删除失败",
				"id":      id,
			})
		}
	})
}

// 根据权限表 构建权限树
