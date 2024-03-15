package middleware

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"fmt"
	"time"
)

type LoginSafe struct {
	ErrorNum   int
	ExpiryTime time.Time
}

var LoginSafeMap map[string]*LoginSafe

// 初始化哈希表
func init() {
	LoginSafeMap = make(map[string]*LoginSafe)
}

// 当用户进行非法操作时，记录其ip以及累计非法操作次数
func CheckLoginError(uip string) int {
	CleanupLoginErrorExpiredData()
	loginSafe, ok := LoginSafeMap[uip]
	if !ok {
		var expiryTime time.Duration = time.Hour * 24
		expiry := time.Now().Add(expiryTime)
		LoginSafeMap[uip] = &LoginSafe{
			ErrorNum:   1,
			ExpiryTime: expiry,
		}
	} else {
		num := loginSafe.ErrorNum
		if num >= 0 && num < utils.LoginNum {
			loginSafe.ErrorNum += 1
		} else if num >= utils.LoginNum {
			return errmsg.CantLogin
		}
	}
	return errmsg.SUCCESS
}

// 检查ip是否还能登录
func CheckIP(uip string) int {
	CleanupLoginErrorExpiredData()
	loginSafe, ok := LoginSafeMap[uip]
	if ok {
		if loginSafe.ErrorNum >= utils.LoginNum {
			return errmsg.CantLogin
		}
	}
	return errmsg.SUCCESS
}

// 清除过期的记录
func CleanupLoginErrorExpiredData() {
	for key, data := range LoginSafeMap {
		if time.Now().After(data.ExpiryTime) {
			delete(LoginSafeMap, key)
		}
	}
}

// 删除记录
func DeleteLoginErrorData(uip string) int {
	fmt.Println(LoginSafeMap[uip])
	delete(LoginSafeMap, uip)
	fmt.Println(LoginSafeMap[uip])
	return errmsg.SUCCESS
}
