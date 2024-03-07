package middleware

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"github.com/pquerna/otp/totp"
)

// two-factor authentication

func GenerateTwoFA(uid string) (string, int) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NyaLog",
		AccountName: uid,
	})
	if err != nil {
		return "", errmsg.ERROR
	}
	return key.Secret(), errmsg.SUCCESS
}
