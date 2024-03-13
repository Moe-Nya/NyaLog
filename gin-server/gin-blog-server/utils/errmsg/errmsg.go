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

	// 管理员用户
	UserInfoError      = 2001
	UserExist          = 2002
	UserNotExist       = 2003
	UserNotValidate    = 2004
	CheckUserFailed    = 2005
	UserEmailCodeError = 2006
	LoginError         = 2007
	UidError           = 2008
	PasswordError      = 2009
	CantLogin          = 2010

	// 邮件发送
	SendEmailFailed = 5001
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
	UserInfoError:      "user information error",
	UserExist:          "user exist",
	UserNotExist:       "user not exist",
	UserNotValidate:    "user not validate",
	CheckUserFailed:    "check user failed",
	UserEmailCodeError: "user email validate code error",
	LoginError:         "an error occurred while logging in",
	UidError:           "uid error",
	PasswordError:      "password error",
	CantLogin:          "this ipdress can't login agin",

	// 邮件发送
	SendEmailFailed: "Email send failed",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
