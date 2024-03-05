package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Createemail(t *testing.T) {
	model.InitDb()
	var email = &model.EmailServer{}
	email.Stmphost = "127.0.0.1"
	email.Stmpport = 110
	email.Emailusername = "aaa"
	email.Emailpassword = "bbb"
	fmt.Println(model.CreateEmai(email))
}
