package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Pw(t *testing.T) {
	fmt.Println(model.ScryptPw("123456"))
}
