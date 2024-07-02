package utils

import (
	"github.com/gin-gonic/gin"
)

// 获取分页参数
type Pagination struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}
type Query struct {
	limit  int
	offset int
}

func GetPagationQuery(ctx *gin.Context) (Query, error) {
	var pagination Pagination
	err := ctx.ShouldBindQuery(&pagination)
	limit := pagination.PageSize
	offset := limit * (limit - 1)
	var query Query
	query.limit = limit
	query.offset = offset
	return query, err
}
