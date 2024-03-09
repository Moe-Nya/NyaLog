package middleware

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"github.com/pquerna/otp/totp"
)

// two-factor authentication

// 生成2FA
func GenerateTwoFA(uid string) (string, string, int) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NyaLog",
		AccountName: uid,
	})
	if err != nil {
		return "", "", errmsg.GenerateSecretFailed
	}
	return key.Secret(), key.URL(), errmsg.SUCCESS
}

// 验证2FA
func Validate2FA(code string, secret string) int {
	validate := totp.Validate(code, secret)
	if validate {
		return errmsg.SUCCESS
	}
	return errmsg.CodeError
}
