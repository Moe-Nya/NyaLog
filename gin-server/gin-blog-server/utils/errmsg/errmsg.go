package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// JWT中间件
	TokenNotExist       = 1001
	TokenInvalid        = 1002
	TokenParseFailed    = 1003
	TokenGenerateFailed = 1004

	// 2FA
	GenerateSecretFailed = 3001
	CodeError            = 3002

	// QRcode
	GenerateQRFailed = 4001

	// 用户登录
	UserInfoError   = 2001
	UserExist       = 2002
	UserNotExist    = 2003
	CheckUserFailed = 2004
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",

	// JWT中间件
	TokenNotExist:       "token not exist",
	TokenInvalid:        "token invalid",
	TokenParseFailed:    "failed to parse claims",
	TokenGenerateFailed: "token generate failed",

	// 2FA
	GenerateSecretFailed: "generate secret failed",
	CodeError:            "code error",

	// QRcode
	GenerateQRFailed: "generate qrcode failed",

	// 管理员用户
	UserInfoError:   "user information error",
	UserExist:       "user exist",
	UserNotExist:    "user not exist",
	CheckUserFailed: "check user failed",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
