package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_NewUser(t *testing.T) {
	model.InitDb()
	user := &model.User{} // 或者 user := new(model.User)
	user.Uid = "moenya"
	user.Username = "moenya"
	user.Password = "123456"
	user.Email = "i@moenya.cat"
	user.Profilephoto = "www.baidu.com"
	user.Slogan = "写代码是热爱，写到世界充满爱"
	user.Secret = "123456"
	fmt.Println(model.NewUser(user))
}
