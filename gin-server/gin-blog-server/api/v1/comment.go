package v1

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户权限
func GetUserOauth(c *gin.Context) {
	code := c.Query("code")
	oauthCode := middleware.GetOauthCode(code)
	token, err := middleware.GetGitHubToken(oauthCode)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "login success",
			"token":   token,
		})
	}
}

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	var userinfo service.CommentUser
	userinfo, err := service.QueryCommentUserInfo(c.Request.Header.Get("Authorization"))
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"message":  "user info get success",
			"userinfo": userinfo,
		})
	}
}

// 新增评论
func NewComment(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "get token failed",
		})
	}
	var data service.Comment
	_ = c.ShouldBindJSON(&data)
	err := service.NewComment(&data, token)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "send comment success",
		})
	}
}

// 读取评论列表(某篇文章)
func SeleCom(c *gin.Context) {
	var comments []model.Comment
	var commentData service.ArticleComments
	_ = c.ShouldBindJSON(&commentData)
	comments, err := service.SeleCom(&commentData)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":      err,
			"articleid": commentData.Articleid,
			"comments":  comments,
		})
	}
}

// 读取评论列表(所有文章)
func SeleAllCom(c *gin.Context) {
	var comments []model.Comment
	var commentData service.ArticleComments
	_ = c.ShouldBindJSON(&commentData)
	comments, err, total := service.SeleAllCom(&commentData)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"total":    total,
			"comments": comments,
		})
	}
}

// 删除评论
func DeleteCom(c *gin.Context) {
	var commentData service.ArticleComments
	_ = c.ShouldBindJSON(&commentData)
	err := service.DeleteCom(&commentData)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "delete comment success",
		})
	}
}
