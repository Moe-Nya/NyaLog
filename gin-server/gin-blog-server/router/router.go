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
	routerv1.Use(middleware.Cors())

	normalrouter := routerv1.Group("api/v1")
	// 用户管理
	// 用户管理-用户是否存在
	normalrouter.GET("/adminexist", v1.UserExist)
	// 用户管理-用户是否验证
	normalrouter.GET("/adminvalidate", v1.UserValidate)
	// 用户管理-注册用户
	normalrouter.POST("/adminregistration", v1.CreateUser)
	// 用户管理-用户登录
	normalrouter.POST("/login", v1.UserLogin)
	// 用户管理-重置密码-用户验证
	normalrouter.POST("/resetpwd", v1.ValidateReset)
	// 用户管理-查询用户公开信息
	normalrouter.GET("/queryuser", v1.QueryUser)

	// 查询博客设置公开信息
	normalrouter.GET("/queryblogset", v1.QueryBlogSet)

	// 文章设置
	// 文章设置-查询单篇文章
	normalrouter.GET("/article/:id", v1.SeleOneArticle)
	// 文章设置-查询文章列表
	normalrouter.POST("/articlelist", v1.SeleListArticle)
	// 文章设置-文章喜欢
	normalrouter.POST("/articlelike", v1.AddLike)

	// 文章分类管理
	// 文章分类管理-查询文章分类列表
	normalrouter.GET("/category", v1.SeleCategory)
	// 文章分类管理-使用cid查询cid名称
	normalrouter.POST("/categoryname", v1.SeleCatName)

	// 页面管理
	// 页面管理-查询页面
	normalrouter.GET("/:purl", v1.SelePage)

	// 评论
	// 评论登录-获取用户授权码
	normalrouter.GET("/githuboauthcode", v1.GetUserOauth)
	// 评论登录-权限控制
	commentUserAuth := routerv1.Group("api/v1/comment")
	commentUserAuth.Use(middleware.CommentToken())
	{
		// 评论-获取用户信息
		commentUserAuth.GET("/comuserinfo", v1.GetUserInfo)
		// 评论-新增评论
		commentUserAuth.POST("/newcomment", v1.NewComment)
	}
	// 评论-读取某篇文章的评论
	normalrouter.POST("/comments", v1.SeleCom)

	// Findme管理
	// Findme管理-查询Findme
	normalrouter.GET("/findmes", v1.SeleFindme)

	// 导航栏管理
	// 导航栏管理-查询所有导航栏列表
	normalrouter.GET("/navigations", v1.SeleNav)

	// 友链管理
	// 友链管理-查询友链
	normalrouter.POST("/friendlinks", v1.SeleFriendLink)

	// 管理员用户权限控制
	authrouter := routerv1.Group("api/v1/" + utils.AdminEntrance)
	authrouter.Use(middleware.JwtToken())
	{
		// token有效性验证
		authrouter.GET("/", v1.Validate)

		// 发送邮件验证码
		authrouter.POST("/sendemail", v1.SendEmailCode)

		// 用户管理
		// 用户管理-验证用户注册信息
		authrouter.POST("/adminvalidate", v1.ValidateUser)
		// 用户管理-登录注销
		authrouter.POST("/loginout", v1.UserLoginOut)
		// 用户管理-重置密码-邮箱验证
		authrouter.POST("/resetpwdcode", v1.ValidateEmail)
		// 用户管理-更改用户信息
		authrouter.POST("/modifyuser", v1.ModifyUser)

		// 博客设置
		// 更新设置
		authrouter.POST("/blogset", v1.UpdateBlogSet)
		// 读取设置
		authrouter.GET("/blogsetlist", v1.SeleBlogSet)

		// 文章设置
		// 文章设置-新增文章
		authrouter.POST("/newarticle", v1.CreateArticle)
		// 文章设置-编辑文章
		authrouter.POST("/editarticle", v1.EditArticle)
		// 文章设置-删除文章
		authrouter.POST("/deletearticle", v1.DeleteArticle)
		// 文章设置-删除同CID文章
		authrouter.POST("/deletecid", v1.DeleCid)

		// 文章分类管理
		// 文章分类管理-新增文章分类
		authrouter.POST("/newcategory", v1.CreateCategory)
		// 文章分类管理-编辑文章分类
		authrouter.POST("/editcategory", v1.EditCategory)
		// 文章分类管理-删除文章分类
		authrouter.POST("/deletecategory", v1.DeleteCategory)

		// 页面管理
		// 页面管理-创建页面
		authrouter.POST("/newpage", v1.CreatePage)
		// 页面管理-查询页面列表
		authrouter.POST("/pages", v1.SelectPageList)
		// 页面管理-编辑页面
		authrouter.POST("/editpage", v1.EditPage)
		// 页面管理-删除页面
		authrouter.POST("/deletepage", v1.DeletePage)

		// 评论
		// 评论-读取所有评论
		authrouter.POST("/allcomments", v1.SeleAllCom)
		// 评论-删除评论
		authrouter.POST("/deletecomments", v1.DeleteCom)

		// Findme管理
		// Findme管理-新增Findme
		authrouter.POST("/newfindme", v1.CreateFindme)
		// Findme管理-修改Findme
		authrouter.POST("/modifyfindme", v1.ModifyFindme)
		// Findme管理-删除Findme
		authrouter.POST("/deletefindme", v1.DeleFindme)

		// 导航栏管理
		// 导航栏管理-新增导航栏
		authrouter.POST("/newnavigation", v1.CreateNav)
		// 导航栏管理-修改导航栏
		authrouter.POST("/editnavigation", v1.ModifyNav)
		// 导航栏管理-删除导航栏
		authrouter.POST("/deletenavigation", v1.DeleNav)

		// 友链管理
		// 友链管理-增加友链
		authrouter.POST("/newfriendlink", v1.CreateFriendLink)
		// 友链管理-修改友链
		authrouter.POST("/modifyfriendlink", v1.ModifyFriendLink)
		// 友链管理-删除友链
		authrouter.POST("/deletefriendlink", v1.DeleFriendLink)
	}

	_ = routerv1.Run(utils.HttpPort)
}
