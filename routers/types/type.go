package types

// ApiError 接口统一返回类型
type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type ApiResponse struct {
	Data  any       `json:"data"`
	Error *ApiError `json:"error"`
}
