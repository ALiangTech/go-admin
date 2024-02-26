package cerrors

const (
	ErrEmptyCredentials = 1
)

func StatusText(code int) string {
	switch code {
	case ErrEmptyCredentials:
		return "账号或者密码不能为空"
	default:
		return "未知错误"
	}
}
