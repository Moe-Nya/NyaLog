package middleware

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 使用用户名来生成token
type UserClaims struct {
	Uid string `json:"uid"`
	Uip string `json:"uip"`
	jwt.RegisteredClaims
}

// 定义JWT密匙
var jwtkey = []byte(utils.JwtKey)

// 创建jwt
func GenerateJWT(user *UserClaims, t time.Duration) (string, int) {
	// 设置过期时间
	expirationTime := time.Now().Add(24 * t)

	claims := &UserClaims{
		user.Uid,
		user.Uip,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", errmsg.TokenGenerateFailed
	}
	if UserLogin(tokenString, user.Uid) != errmsg.SUCCESS {
		return "", errmsg.TokenGenerateFailed
	}
	return tokenString, errmsg.SUCCESS
}

// 验证token
func ValidateJWT(tokenString string, userip string) (string, int) {
	// 解析jwt
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})
	if err != nil {
		return "", errmsg.TokenParseFailed
	}
	// 验证token有效性
	if !token.Valid {
		return "", errmsg.TokenInvalid
	}
	uid, ok := token.Claims.(*UserClaims)
	if !ok {
		return "", errmsg.TokenParseFailed
	}
	if tokenString == GetToken(uid.Uid) || uid.Uip == userip {
		return uid.Uid, errmsg.SUCCESS
	}
	return "", errmsg.ERROR
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		// 接收传入的token
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			code = errmsg.TokenNotExist
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"info":   errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 验证token
		_, code = ValidateJWT(tokenString, c.ClientIP())
		if code != errmsg.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"info":   errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
