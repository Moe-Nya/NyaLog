package middleware

import "NyaLog/gin-blog-server/utils/errmsg"

var UserTokenMap map[string]string

// 初始化usertoken哈希表
func init() {
	UserTokenMap = make(map[string]string)
}

// 登陆时存入token和对应的uid
func UserLogin(token string, uid string) int {
	UserTokenMap[token] = uid
	return errmsg.SUCCESS
}

// 提取token对应的uid
func GetUid(token string) string {
	return UserTokenMap[token]
}
