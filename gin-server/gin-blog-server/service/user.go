package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 创建用户
// todo 先解决谷歌二级验证码中间件、登录安全中间件、token内存保持服务
func CreateUser(user *model.User) int {
	if user.Uid == "" || user.Username == "" || user.Password == "" || user.Email == "" {
		return errmsg.UserInfoError
	}

	return errmsg.SUCCESS
}
