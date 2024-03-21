package cerrors

const (
	ErrEmptyCredentials = 1
	ErrUser             = 2
	ErrGenJwt           = 3
)

func StatusText(code int) string {
	switch code {
	case ErrEmptyCredentials:
		return "账号或者密码不能为空"
	case ErrUser:
		return "账号不存在或者密码错误"
	case ErrGenJwt:
		return "jwt认证生成失败"
	default:
		return "未知错误"
	}
}
