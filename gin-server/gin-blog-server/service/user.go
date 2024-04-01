package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"math/rand"
	"regexp"
	"time"
)

// 判断用户是否存在
func UserExist() (bool, int) {
	return model.UserExist()
}

// 判断用户是否验证
func UserValidate() (bool, int) {
	return model.UserValidate()
}

// -------------------------------------------

// 发送邮箱时需要的结构体
type SendEmail struct {
	Uid      string `json:"uid"`
	Email    string `json:"email"`
	Useplace string `json:"useplace"`
	Category bool   `json:"category"`
}

// 发送邮件验证码
// 传值false为注册时验证用户， true为找回密码所用
func SendEmailCode(data *SendEmail) int {
	// 验证用户是否已经存在
	userexist, err := UserExist()
	if err == errmsg.ERROR {
		return err
	}
	if userexist {
		uservalidate, err := model.UserValidate()
		if err == errmsg.ERROR {
			return err
		}
		if uservalidate && !data.Category {
			return errmsg.UserExist
		}
	} else {
		return errmsg.UserNotExist
	}

	var user model.User
	user, err = model.SeleUser()
	if err == errmsg.ERROR {
		return err
	}
	if data.Email == user.Email {
		// 生成邮件验证码并且发送邮件信息
		rand.NewSource(time.Now().UnixNano())
		code := make([]byte, 6)
		for i := 0; i < 6; i++ {
			code[i] = byte(rand.Intn(10) + '0')
		}
		// 编辑发送邮件的信息
		msg := []byte("From: " + "NyaLog" + "\r\n" +
			"To: " + user.Email + "\r\n" +
			"Subject: " + "NyaLog: Your verification code is:" + "\r\n" +
			"\r\n" +
			string(code))
		err = middleware.SendEmail(user.Email, msg)
		if err != errmsg.SUCCESS {
			return err
		}
		middleware.CheckUserEmailCode(user.Uid, data.Useplace, string(code))
	} else {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// -------------------------------------------

// 创建用户
func CreateUser(user *model.User) ([]byte, int) {
	// 验证用户是否已经存在
	userexist, err := UserExist()
	if err == errmsg.ERROR {
		return nil, errmsg.ERROR
	}
	// 若存在则判断是否验证，若验证则返回用户存在，若未验证则删除该用户重新注册
	if userexist {
		uservalidate, err := model.UserValidate()
		if err == errmsg.ERROR {
			return nil, errmsg.ERROR
		}
		if uservalidate {
			return nil, errmsg.UserExist
		} else {
			model.DeleteUser()
		}
	}
	// 正则判断密码是否含有数字、大小写字母、标点符号；密码长度需要大于6
	re := regexp.MustCompile(`[0-9]+[a-zA-Z]+[!@#$%^&*().]+`)
	match := re.MatchString(user.Password)
	emaiRe := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emaiMatch := emaiRe.MatchString(user.Email)

	if user.Uid == "" || user.Username == "" || !match || !emaiMatch || len(user.Password) <= utils.PasswordMinLen || user.Email == "" {
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

// 注册用户时验证用户需要用到的结构体
type CheckUserToken struct {
	Emailcode string `json:"emailcode"`
	Twofacode string `json:"twofacode"`
	Userip    string `json:"userip"`
}

// 创建用户时的验证
func CheckUser(data *CheckUserToken, token string) int {
	// 验证用户是否已经存在
	userexist, err := UserExist()
	if err == errmsg.ERROR {
		return errmsg.ERROR
	}
	if userexist {
		uservalidate, err := model.UserValidate()
		if err == errmsg.ERROR {
			return errmsg.ERROR
		}
		if uservalidate {
			return errmsg.UserExist
		}
	} else {
		return errmsg.UserNotExist
	}

	var user model.User
	user, _ = model.SeleUser()

	// 验证用户输入的验证信息
	uid, _ := middleware.ValidateJWT(token, data.Userip)
	code, useplace, err := middleware.GetCode(uid)
	if err != errmsg.SUCCESS || code != data.Emailcode {
		return errmsg.ValidateCodeError
	}
	if useplace != "create" {
		return errmsg.ValidateCodeError
	}
	if middleware.Validate2FA(data.Twofacode, user.Secret) == errmsg.CodeError {
		return errmsg.CodeError
	}
	user.Validateuser = 1
	err = model.UpdateUser(user.Uid, &user)
	if err == errmsg.ERROR {
		return errmsg.ERROR
	}
	middleware.DeleteCode(uid)
	middleware.DeleteToken(uid)
	return errmsg.SUCCESS
}

// 用户登录时所用的结构体
type Ulogin struct {
	Uid       string `json:"uid"`
	Password  string `json:"password"`
	TwoFACode string `json:"twofacode"`
}

// -------------------------------------------

// 用户登录
func UserLogin(data *Ulogin) int {
	// 验证用户是否已经存在
	userexist, err := UserExist()
	if err == errmsg.ERROR {
		return errmsg.LoginError
	}
	if userexist {
		uservalidate, err := model.UserValidate()
		if err == errmsg.ERROR {
			return errmsg.LoginError
		}
		if !uservalidate {
			return errmsg.UserNotValidate
		}
	} else {
		return errmsg.UserNotExist
	}

	var user model.User
	user, err = model.SeleUser()
	if err != errmsg.SUCCESS {
		return errmsg.LoginError
	}

	// 2FA验证
	twoFAValidate := middleware.Validate2FA(data.TwoFACode, user.Secret)
	if twoFAValidate != errmsg.SUCCESS {
		return twoFAValidate
	}

	// Uid验证
	if data.Uid != user.Uid {
		return errmsg.UidError
	}

	// Password验证
	pwValidate := model.VertifyPw(data.Password)
	if pwValidate != errmsg.SUCCESS {
		return errmsg.PasswordError
	}
	return errmsg.SUCCESS
}

// 登录注销
func UserLoginOut(uid string) int {
	err := middleware.DeleteToken(uid)
	if err != errmsg.SUCCESS {
		return errmsg.TokenNotExist
	}
	return errmsg.SUCCESS
}

// -------------------------------------------

// 重置密码时所用的结构体
type ValidateRe struct {
	Uid       string `json:"uid"`
	Email     string `json:"email"`
	TwoFaCode string `json:"twofacode"`
}

// 重置密码
func ValidateReset(data *ValidateRe) int {
	var user model.User
	user, err := model.SeleUser()
	if err != errmsg.SUCCESS {
		return errmsg.ResetPasswordFailed
	}
	if data.Uid != user.Uid || data.Email != user.Email {
		return errmsg.UserInfoError
	} else {
		err = middleware.Validate2FA(data.TwoFaCode, user.Secret)
		if err != errmsg.SUCCESS {
			return err
		}
	}
	return errmsg.SUCCESS
}

// 重置密码验证邮箱所用结构体
type ValidateEm struct {
	Uid      string `json:"uid"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

// 验证邮箱
func ValidateEmail(data *ValidateEm) int {
	// 正则判断密码是否含有数字、大小写字母、标点符号；密码长度需要大于6
	re := regexp.MustCompile(`[0-9]+[a-zA-Z]+[!@#$%^&*().]+`)
	match := re.MatchString(data.Password)
	if !match {
		return errmsg.ResetPasswordFailed
	}
	code, useplace, err := middleware.GetCode(data.Uid)
	if err != errmsg.SUCCESS {
		return errmsg.ValidateCodeError
	}
	if code != data.Code {
		return errmsg.ValidateCodeError
	}
	if useplace != "reset" {
		return errmsg.ValidateCodeError
	}
	err = model.ModifyPw(data.Password)
	if err != errmsg.SUCCESS {
		return errmsg.ResetPasswordFailed
	}
	return errmsg.SUCCESS
}

// -------------------------------------------
// 公开用户信息结构体
type UserInfo struct {
	Username     string `json:"username"`
	Profilephoto string `json:"profilephoto"`
	Email        string `json:"email"`
	Slogan       string `json:"slogan"`
}

// 查询用户可公开信息
func QueryUser() (UserInfo, int) {
	var user model.User
	var userinfo UserInfo
	user, err := model.QueryUser()
	if err != errmsg.SUCCESS {
		return userinfo, errmsg.QueryUserFailed
	}
	userinfo.Username = user.Username
	userinfo.Profilephoto = user.Profilephoto
	userinfo.Email = user.Email
	userinfo.Slogan = user.Slogan
	return userinfo, errmsg.SUCCESS
}

// 更改用户可公开信息
func ModifyUser(data *UserInfo) int {
	var user model.User
	user, err := model.SeleUser()
	if err != errmsg.SUCCESS {
		return errmsg.ModifyUserInfoFailed
	}
	if data.Username != "" {
		user.Username = data.Username
	}
	if data.Profilephoto != "" {
		user.Profilephoto = data.Profilephoto
	}
	emaiRe := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	emaiMatch := emaiRe.MatchString(data.Email)
	if !emaiMatch {
		return errmsg.ModifyUserInfoFailed
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.Slogan != "" {
		user.Slogan = data.Slogan
	}
	err = model.UpdateUser(user.Uid, &user)
	if err != errmsg.SUCCESS {
		return errmsg.ModifyUserInfoFailed
	}
	return errmsg.SUCCESS
}
