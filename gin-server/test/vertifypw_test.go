package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_VertifyPw(t *testing.T) {
	model.InitDb()
	fmt.Println(model.VertifyPw("123456789"))
}
