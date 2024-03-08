package main

import (
	"NyaLog/gin-blog-server/middleware"
	"fmt"
	"testing"
)

func TestGenerateJwt(t *testing.T) {
	var user = &middleware.UserClaims{}
	user.Uid = "moenya"
	user.Uip = "127.0.0.1"
	token, _ := middleware.GenerateJWT(user)

	fmt.Println(middleware.ValidateJWT(token, user.Uip))
}
