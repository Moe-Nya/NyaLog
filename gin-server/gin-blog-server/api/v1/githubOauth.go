package v1

import (
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取用户权限
func GetUserOauth(c *gin.Context) {
	code := c.Query("code")
	oauthCode := service.GetOauthCode(code)
	token, err := service.GetToken(oauthCode)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		userinfo, err := service.GetUserInfo(token)
		if err != errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":     err,
				"message":  "login success",
				"userinfo": userinfo,
			})
		}
	}

}
