package routes

import (
	"aliangtect/go-admin/db"
	"aliangtect/go-admin/routers/types"
	"aliangtect/go-admin/routers/v1/middlewares"
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

type Role struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Policy      []string `json:"policy"`
}

// 向数据库插入数据

func interRole(role Role) (bool, error) {
	//db.DB.Exec("INSERT INTO roles (name,description,policy) values (?, ?, ?)", role.Name, role.Description, pq.Array(role.Policy))
	// 通过casbin api 向casbin_rule 插入一个角色
	// AddGroupingPolicies()
	e := middlewares.Enforcer
	return e.AddRolesForUser(role.Name, role.Policy)
}

// 删除角色

func delRole(id string) error {
	res := db.DB.Exec("DELETE FROM roles where id = ? ", id)
	return res.Error
}

// 判断角色是否存在
func isExistRoleName(name string) bool {
	res := db.DB.Exec("SELECT * from roles where name = ?", name)
	fmt.Printf("res %+v", res)
	fmt.Println("---")
	return res.RowsAffected > 0 // 如果大于 说明存在相同的角色名称
}

/*
* 当前用户创建角色 得确保当前用户提交的权限code 范围是当前用户权限的子集
* 防止通过接口提交 或者其他方式 让子级账号越 权限
* 1、查询当前账号拥有的权限
* 2、判断创建角色的policy 是否全部合规 并且是当前账号的子集
* 3、提交数据 通过casbin 创建角色
* 4、提示前端创建成功
 */

/*
* 获取角色列表
 */
// 包含ID的结构体
type RoleWithID struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Policy      []string `json:"policy"`
}

func gerRoles() ([]RoleWithID, error) {
	var roles []RoleWithID
	stmt := db.DB.Raw("select * from roles").Scan(&roles)
	return roles, stmt.Error
}

func RetrieveRole(router *gin.RouterGroup) {
	router.GET("/roles", func(ctx *gin.Context) {
		roles, err := gerRoles()
		if err != nil {
			ctx.JSON(200, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Status:     0,
					StatusText: "角色列表获取失败",
				},
			})
		} else {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data:  roles,
				Error: nil,
			})
		}

	})
	// 创建角色
	// 获取角色name desc policy
	router.POST("/role", func(ctx *gin.Context) {
		role := Role{}
		ctx.ShouldBindJSON(&role)
		isExistName := isExistRoleName(role.Name)
		if isExistName {
			// 存在相同角色
			ctx.JSON(200, gin.H{
				"message": "角色创建失败,角色名称不能重复",
				"role":    role,
			})
			return
		}
		err := postRole(role)
		if err != nil {
			ctx.JSON(http.StatusOK, types.ApiResponse{
				Data: nil,
				Error: &types.ApiError{
					Status:     0,
					StatusText: err.Error(),
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, types.ApiResponse{
			Data:  nil,
			Error: nil,
		})
		// if res != nil {
		// 	// 获取角色name 权限policy
		// 	ctx.JSON(200, gin.H{
		// 		"message": "角色创建失败",
		// 		"role":    role,
		// 	})
		// } else {
		// 	// 获取角色name 权限policy
		// 	ctx.JSON(200, gin.H{
		// 		"message": "角色创建成功",
		// 		"role":    role,
		// 		"result":  res,
		// 	})
		// }
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

/*
@desc 创建角色
通过事务方式：一个操作失败了 数据会回滚
*/
func postRole(role Role) error {
	rules := [][]string{}
	e := middlewares.Enforcer
	// e.HasPermissionForUser
	for _, code := range role.Policy {
		rules = append(rules, []string{role.Name, code}) // todo code 必须是属于创建人的 否则属于恶意创建角色，则取消创建
	}

	// 开始事务
	tx := db.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	// 插入角色到数据
	result := tx.Exec("INSERT INTO roles (name, description) values(?,?)", role.Name, role.Description)
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("插入角色失败：%w", result.Error)
	}

	// 添加casbin 角色
	added, err := e.AddGroupingPolicies(rules) // todo 回滚机制 如果INSERT 执行失败 则 Add操作需要回滚
	fmt.Println(added)
	if err != nil || !added { // 需要added 判断是为了判断 防止rules 里面的权限提交一下不存在的权限code
		tx.Rollback()
		return fmt.Errorf("casbin 角色添加失败: %w。added:%t", err, added)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		// 回滚Casbin策略
		if _, err := e.RemoveGroupingPolicies(rules); err != nil {
			return fmt.Errorf("casbin 角色回滚失败: %w", err)
		}
		return err
	}
	return nil
}

// 检查权限 permission 是否属于当前角色创建人
func checkPermission(enforcer *casbin.Enforcer, role, permission []string) (bool, error) {
	// 检查角色是否拥有权限
	return enforcer.HasPermissionForUser(role, permission)
}
