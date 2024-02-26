package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Pw(t *testing.T) {
	pw, _ := model.ScryptPw("123456")
	fmt.Println(pw)
}
