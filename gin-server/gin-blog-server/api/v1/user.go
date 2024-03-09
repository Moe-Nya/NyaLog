package v1

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 判断用户是否存在
func UserExist(c *gin.Context) {
	userexist, err := service.UserExist()
	if err == errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "User system error",
		})
	}
	if userexist {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "User exist",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": errmsg.SUCCESS,
		})
	}
}

// 注册用户
func CreateUser(c *gin.Context) {
	var data model.User
	var userip = c.ClientIP()
	_ = c.ShouldBindJSON(&data)
	qrcode, err := service.CreateUser(&data)
	if err != errmsg.SUCCESS {
		if err == errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR,
				"message": "User registration failed",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		}
	} else {
		var theTime time.Duration = time.Minute
		var userclaims middleware.UserClaims
		userclaims.Uid = data.Uid
		userclaims.Uip = userip
		token, err := middleware.GenerateJWT(&userclaims, theTime)
		if err == errmsg.TokenGenerateFailed {
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.TokenGenerateFailed,
				"message": "Token generate error, registration failed",
			})
		}
		middleware.UserLogin(data.Uid, token)
		c.JSON(http.StatusOK, gin.H{
			"code":      errmsg.SUCCESS,
			"message":   "User registration success",
			"temptoken": token,
			"qrcode":    qrcode,
		})
	}
}

// 发送邮件验证码
func SendEmailCode(c *gin.Context) {
	err := service.SendEmailCode()
	if err != errmsg.SUCCESS {
		if err == errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": "Send code failed",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "Send success",
		})
	}
}

// 验证注册的用户
func ValidateUser(c *gin.Context) {
	var data service.CheckUserToken
	_ = c.ShouldBindJSON(&data)
	data.Userip = c.ClientIP()
	tokenString := c.Request.Header.Get("Authorization")
	err := service.CheckUser(&data, tokenString)
	if err != errmsg.SUCCESS {
		if err == errmsg.ERROR {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": "User validate error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "Validate success",
		})
	}
}

// 用户登录
