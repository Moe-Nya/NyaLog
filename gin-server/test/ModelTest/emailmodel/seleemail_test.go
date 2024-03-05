package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_ReadEmail(t *testing.T) {
	model.InitDb()
	var email model.EmailServer
	email, _ = model.SeleEmail()
	fmt.Println(email)
}
