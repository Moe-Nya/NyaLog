package main

import (
	"NyaLog/gin-blog-server/middleware"
	"fmt"
	"testing"
	"time"
)

func TestGenerateJwt(t *testing.T) {
	var user = &middleware.UserClaims{}
	user.Uid = "moenya"
	user.Uip = "127.0.0.1"
	var tt time.Duration = time.Hour
	token, _ := middleware.GenerateJWT(user, tt)

	fmt.Println(middleware.ValidateJWT(token, user.Uip))
}
