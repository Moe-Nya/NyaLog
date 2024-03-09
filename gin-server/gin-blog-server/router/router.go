package router

import (
	v1 "NyaLog/gin-blog-server/api/v1"
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/utils"

	"github.com/gin-gonic/gin"
)

func IniRouter() {
	gin.SetMode(utils.AppMode)
	routerv1 := gin.Default()

	normalrouter := routerv1.Group("api/v1")
	// 用户是否存在
	normalrouter.GET("/adminexist", v1.UserExist)
	normalrouter.POST("/adminregistration", v1.CreateUser)

	authrouter := routerv1.Group("api/v1")
	authrouter.Use(middleware.JwtToken())
	{
		// 发送邮件验证码
		authrouter.GET("/sendemail", v1.SendEmailCode)
		// 验证用户注册信息
		authrouter.POST("/adminvalidate", v1.ValidateUser)
	}

	_ = routerv1.Run(utils.HttpPort)
}
