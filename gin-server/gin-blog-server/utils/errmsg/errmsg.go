package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// JWT中间件
	TokenNotExist    = 1001
	TokenInvalid     = 1002
	TokenParseFailed = 1003
)

var codeMsg = map[int]string{
	SUCCESS:          "ok",
	ERROR:            "error",
	TokenNotExist:    "token not exist",
	TokenInvalid:     "token invalid",
	TokenParseFailed: "failed to parse claims",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
