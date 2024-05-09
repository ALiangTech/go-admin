package types

// ApiError 接口统一返回类型
type ApiError struct {
	Status     int    `json:"status"`
	StatusText string `json:"statusText"`
}
type ApiResponse struct {
	Data  any       `json:"data"`
	Error *ApiError `json:"error"`
}

// 分页查询 统一

type ApiQueryRequest struct {
	Page int `json:"page"` // 页码
	Size int `json:"size"` // 一页数据展示多少数量
}

// 分页统一返回格式

type Pagation[T any] struct {
	Record    []T `json:"record"`
	PageCount int `json:"pageCount"`
}
