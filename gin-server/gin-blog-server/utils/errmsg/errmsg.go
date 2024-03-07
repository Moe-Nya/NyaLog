package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// JWT中间件
	TokenNotExist    = 1001
	TokenInvalid     = 1002
	TokenParseFailed = 1003

	// 用户登录
	UserInfoError = 2001
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",

	// JWT中间件
	TokenNotExist:    "token not exist",
	TokenInvalid:     "token invalid",
	TokenParseFailed: "failed to parse claims",
	UserInfoError:    "uid or password error",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
