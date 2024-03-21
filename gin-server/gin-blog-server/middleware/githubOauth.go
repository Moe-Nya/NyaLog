package middleware

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接收授权码
func GetOauthCode(code string) string {
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		utils.GitHubID, utils.GitHUbSecret, code,
	)
}

// Token
type Token struct {
	Token string `json:"access_token"`
}

// 获取Token
func GetGitHubToken(url string) (string, int) {
	// 形成请求
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, url, nil); err != nil {
		return "", errmsg.CreateQueryFailed
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		return "", errmsg.GetQueryFailed
	}
	var token Token
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		return "", errmsg.GetQueryFailed
	}
	return token.Token, errmsg.SUCCESS
}

// 获取用户请求
func GetUserInfo(token string) (map[string]interface{}, int) {
	// 形成请求
	var userInfoUrl = "https://api.github.com/user" // github用户信息获取接口
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, userInfoUrl, nil); err != nil {
		return nil, errmsg.CreateQueryFailed
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 发送请求并获取响应
	var client = http.Client{}
	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return nil, errmsg.GetQueryFailed
	}

	// 将响应的数据写入 userInfo 中，并返回
	var userInfo = make(map[string]interface{})
	if err = json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return nil, errmsg.GetQueryFailed
	}
	return userInfo, errmsg.SUCCESS
}

// 回复评论中间件
func CommentToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			code := errmsg.TokenNotExist
			c.JSON(http.StatusOK, gin.H{
				"status": code,
				"info":   errmsg.GetErrorMsg(code),
			})
			c.Abort()
			return
		}
		// 验证token
		_, code := GetUserInfo(tokenString)
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

// todo--1.优化代码（建立中间件）2.写好评论功能 3.建立friend link model层 4.可以做nav、findme、friend link的api
