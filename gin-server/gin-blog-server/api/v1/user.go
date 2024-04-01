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
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "User system error",
		})
	} else {
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
}

// 判断用户是否验证
func UserValidate(c *gin.Context) {
	uservalidate, err := service.UserValidate()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ERROR,
			"message": "User system error",
		})
	} else {
		if uservalidate {
			c.JSON(http.StatusOK, gin.H{
				"code":    errmsg.ERROR,
				"message": "User validated",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": errmsg.SUCCESS,
			})
		}
	}
}

// -------------------------------------------

// 发送邮件验证码
func SendEmailCode(c *gin.Context) {
	var data service.SendEmail
	_ = c.ShouldBindJSON(&data)
	if middleware.CheckIP(c.ClientIP()) != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.SendEmailFailed,
			"message": errmsg.GetErrorMsg(errmsg.SendEmailFailed),
		})
	} else {
		var user model.User
		user, err := model.SeleUser()
		if err != errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": "validate failed",
			})
		} else {
			if user.Validateuser == 1 {
				data.Category = true
				data.Useplace = "reset"
			} else {
				data.Category = false
				data.Useplace = "create"
			}
			err = service.SendEmailCode(&data)
			if err != errmsg.SUCCESS {
				if err == errmsg.ERROR {
					c.JSON(http.StatusOK, gin.H{
						"code":    err,
						"message": "send code failed",
					})
				} else {
					c.JSON(http.StatusOK, gin.H{
						"code":    err,
						"message": errmsg.GetErrorMsg(err),
					})
				}
			} else {
				_ = middleware.CheckLoginError(c.ClientIP())
				c.JSON(http.StatusOK, gin.H{
					"code":    err,
					"message": "send success",
				})
			}
		}
	}
}

// -------------------------------------------

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
		middleware.DeleteLoginErrorData(c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "Validate success",
		})
	}
}

// -------------------------------------------

// 用户登录
func UserLogin(c *gin.Context) {
	err := middleware.CheckIP(c.ClientIP())
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		var data service.Ulogin
		_ = c.ShouldBindJSON(&data)
		err = service.UserLogin(&data)
		if err != errmsg.SUCCESS {
			// 用户登录错误次数
			e := middleware.CheckLoginError(c.ClientIP())
			if e != errmsg.SUCCESS {
				c.JSON(http.StatusOK, gin.H{
					"code":    e,
					"message": errmsg.GetErrorMsg(e),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		} else {
			var user middleware.UserClaims
			user.Uid = data.Uid
			user.Uip = c.ClientIP()
			var t time.Duration = time.Hour
			tokenString, err := middleware.GenerateJWT(&user, t)
			if err != errmsg.SUCCESS {
				c.JSON(http.StatusOK, gin.H{
					"code":    err,
					"message": errmsg.GetErrorMsg(err),
				})
			} else {
				middleware.UserLogin(data.Uid, tokenString)
				middleware.DeleteLoginErrorData(c.ClientIP())
				c.JSON(http.StatusOK, gin.H{
					"code":    err,
					"message": errmsg.GetErrorMsg(err),
					"token":   tokenString,
				})
			}
		}
	}
}

// 登录注销
func UserLoginOut(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	uid, err := middleware.ValidateJWT(tokenString, c.ClientIP())
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		err = service.UserLoginOut(uid)
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	}
}

// -------------------------------------------

// 重置密码
func ValidateReset(c *gin.Context) {
	err := middleware.CheckIP(c.ClientIP())
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		var validatereset service.ValidateRe
		_ = c.ShouldBindJSON(&validatereset)
		err = service.ValidateReset(&validatereset)
		if err != errmsg.SUCCESS {
			middleware.CheckLoginError(c.ClientIP())
			c.JSON(http.StatusOK, gin.H{
				"code":    err,
				"message": errmsg.GetErrorMsg(err),
			})
		} else {
			var user middleware.UserClaims
			user.Uid = validatereset.Uid
			user.Uip = c.ClientIP()
			var t time.Duration = time.Minute
			tokenString, err := middleware.GenerateJWT(&user, t)
			if middleware.UserLogin(user.Uid, tokenString) != errmsg.SUCCESS {
				c.JSON(http.StatusOK, gin.H{
					"code":    err,
					"message": errmsg.GetErrorMsg(err),
				})
			} else {
				middleware.DeleteLoginErrorData(c.ClientIP())
				c.JSON(http.StatusOK, gin.H{
					"code":    err,
					"message": errmsg.GetErrorMsg(err),
					"token":   tokenString,
				})
			}
		}
	}
}

// 验证邮箱
func ValidateEmail(c *gin.Context) {
	var user service.ValidateEm
	_ = c.ShouldBindJSON(&user)
	err := service.ValidateEmail(&user)
	if err != errmsg.SUCCESS {
		middleware.CheckLoginError(c.ClientIP())
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		middleware.DeleteCode(user.Uid)
		middleware.DeleteToken(user.Uid)
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.SUCCESS,
			"message": "reset success",
		})
	}
}

// -------------------------------------------
// 查询用户可公开信息
func QueryUser(c *gin.Context) {
	var user service.UserInfo
	user, err := service.QueryUser()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"userinfo": user,
		})
	}
}

// 更改用户可公开信息
func ModifyUser(c *gin.Context) {
	var data service.UserInfo
	_ = c.ShouldBindJSON(&data)
	err := service.ModifyUser(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "modify user info success",
		})
	}
}
