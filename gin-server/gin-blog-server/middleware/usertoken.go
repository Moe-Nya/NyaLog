package middleware

import "NyaLog/gin-blog-server/utils/errmsg"

var UserTokenMap map[string]string
var UserEmailCodeMap map[string]string

// 初始化usertoken哈希表
func init() {
	UserTokenMap = make(map[string]string)
	UserEmailCodeMap = make(map[string]string)
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
func UserEmailCode(uid string, code string) int {
	UserEmailCodeMap[uid] = code
	return errmsg.SUCCESS
}

// 提取uid对应的code
func GetCode(uid string) string {
	return UserEmailCodeMap[uid]
}

// 删除uid对应的code
func DeleteCode(uid string) int {
	delete(UserEmailCodeMap, uid)
	return errmsg.SUCCESS
}
