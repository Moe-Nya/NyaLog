package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_SeleUser(t *testing.T) {
	model.InitDb()
	user, _ := model.SeleUser()
	fmt.Println(111)
	fmt.Println(user.Username)
}
