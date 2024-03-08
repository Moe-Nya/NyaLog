package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"encoding/base64"
	"net/http"

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
	_ = c.ShouldBindJSON(&data)
	qrcode, err := service.CreateUser(&data)
	if err == errmsg.ERROR {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "User registration failed",
		})
	}
	if err == errmsg.UserExist {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.UserExist,
			"message": errmsg.GetErrorMsg(errmsg.UserExist),
		})
	}
	if err == errmsg.UserInfoError {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.UserInfoError,
			"message": errmsg.GetErrorMsg(errmsg.UserInfoError),
		})
	}
	if err == errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.SUCCESS,
			"message": "User registration success",
			"qrcode":  base64.StdEncoding.EncodeToString(qrcode),
		})
	}
}
