package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Update(t *testing.T) {
	model.InitDb()
	var user model.User
	user, _ = model.SeleUser()
	user.Username = "MoeNya"
	code := model.UpdateUser(user.Uid, user)
	fmt.Println(code)
}
