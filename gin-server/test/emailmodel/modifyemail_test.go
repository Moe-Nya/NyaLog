package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_ModifyEmail(t *testing.T) {
	model.InitDb()
	var email model.EmailServer
	email, _ = model.SeleEmail()
	email.Emailpassword = "ddd"
	fmt.Println(model.ModifyEmail(email))
}
