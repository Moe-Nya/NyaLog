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
	routerv1.RedirectTrailingSlash = false

	normalrouter := routerv1.Group("api/v1")
	// 用户是否存在
	normalrouter.GET("/adminexist", v1.UserExist)

	// 注册用户
	normalrouter.POST("/adminregistration", v1.CreateUser)

	// 用户登录
	normalrouter.POST("/login", v1.UserLogin)

	// 重置密码-用户验证
	normalrouter.POST("/resetpwd", v1.ValidateReset)

	authrouter := routerv1.Group("api/v1/" + utils.AdminEntrance)
	authrouter.Use(middleware.JwtToken())
	{
		// 发送邮件验证码
		authrouter.GET("/sendemail", v1.SendEmailCode)

		// 验证用户注册信息
		authrouter.POST("/adminvalidate", v1.ValidateUser)

		// 登录注销
		authrouter.POST("/loginout", v1.UserLoginOut)

		// 重置密码-邮箱验证
		authrouter.POST("/resetpwdcode", v1.ValidateEmail)

		// 博客设置
		// 更新设置
		authrouter.POST("/blogset", v1.UpdateBlogSet)
		// 读取设置
		authrouter.GET("/blogsetlist", v1.SeleBlogSet)

		// 文章设置
		// 文章设置-新增文章
		authrouter.POST("/newarticle", v1.CreateArticle)
		// 文章设置-查询单篇文章
		authrouter.GET("/article/:id", v1.SeleOneArticle)
		// 文章设置-查询文章列表
		authrouter.GET("/articlelist", v1.SeleListArticle)
	}

	_ = routerv1.Run(utils.HttpPort)
}
