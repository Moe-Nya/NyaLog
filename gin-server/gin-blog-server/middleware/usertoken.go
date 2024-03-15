package middleware

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"time"
)

type UserEmailCode struct {
	Emailcode  string `json:"emailcode"`
	Useplace   string `json:"useplace"`
	ExpiryTime time.Time
}

var UserTokenMap map[string]string
var UserEmailCodeMap map[string]UserEmailCode

// 初始化哈希表
func init() {
	UserTokenMap = make(map[string]string)
	UserEmailCodeMap = make(map[string]UserEmailCode)
}

// 登陆时存入uid和对应的token
func UserLogin(uid string, token string) int {
	UserTokenMap[uid] = token
	return errmsg.SUCCESS
}

// 提取uid对应的token
func GetToken(uid string) string {
	return UserTokenMap[uid]
}

// 删除uid对应的token
func DeleteToken(uid string) int {
	delete(UserTokenMap, uid)
	return errmsg.SUCCESS
}

// 验证时存入用户uid和邮箱验证码
func CheckUserEmailCode(uid string, useplace string, code string) int {
	var expiryTime time.Duration = time.Minute * 30
	expiry := time.Now().Add(expiryTime)
	UserEmailCodeMap[uid] = UserEmailCode{
		Emailcode:  code,
		Useplace:   useplace,
		ExpiryTime: expiry,
	}
	return errmsg.SUCCESS
}

// 提取uid对应的code
func GetCode(uid string) (string, string, int) {
	CleanupEmaiCodeExpiredData()
	data, ok := UserEmailCodeMap[uid]
	if ok && time.Now().Before(data.ExpiryTime) {
		return data.Emailcode, data.Useplace, errmsg.SUCCESS
	}
	return "", "", errmsg.ERROR
}

// 删除uid对应的code
func DeleteCode(uid string) int {
	delete(UserEmailCodeMap, uid)
	return errmsg.SUCCESS
}

// 清除过期emaicode
func CleanupEmaiCodeExpiredData() {
	for key, data := range UserEmailCodeMap {
		if time.Now().After(data.ExpiryTime) {
			delete(UserEmailCodeMap, key)
		}
	}
}
