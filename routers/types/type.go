package types

// 接口统一返回类型
type ApiError struct {
	Code    int
	Message string
}
type ApiResponse struct {
	Data  any
	Error ApiError
}
