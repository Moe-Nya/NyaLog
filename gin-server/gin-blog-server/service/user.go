package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"regexp"
)

// 判断用户是否存在
func UserExist() (bool, int) {
	return model.UserExist()
}

// 创建用户
// todo 登录安全中间件
func CreateUser(user *model.User) ([]byte, int) {
	// 验证用户是否已经存在
	userexist, err := UserExist()
	if err == errmsg.ERROR {
		return nil, errmsg.ERROR
	}
	if userexist {
		return nil, errmsg.UserExist
	}
	// 正则判断密码是否含有数字、大小写字母、标点符号；密码长度需要大于6
	re := regexp.MustCompile(`[0-9]+[a-zA-Z]+[!@#$%^&*().]+`)
	match := re.MatchString(user.Password)
	if user.Uid == "" || user.Username == "" || !match || len(user.Password) <= utils.PasswordMinLen || user.Email == "" {
		return nil, errmsg.UserInfoError
	}

	// 账户信息通过后，开始认证2FA
	secret, url, err := middleware.GenerateTwoFA(user.Uid)
	user.Secret = secret
	if err == errmsg.GenerateSecretFailed {
		return nil, errmsg.GenerateQRFailed
	}
	qrcode, err := middleware.GenerateQRcode(url, 256)
	if err == errmsg.GenerateQRFailed {
		return nil, errmsg.GenerateQRFailed
	}

	// 用户注册
	err = model.NewUser(user)
	if err == errmsg.ERROR {
		return nil, errmsg.ERROR
	}
	return qrcode, errmsg.SUCCESS
}
