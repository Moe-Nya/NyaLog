package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_SeleUser(t *testing.T) {
	model.InitDb()
	user, _ := model.SeleUser()
	fmt.Println(user.Username)
}
