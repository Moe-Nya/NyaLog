package main

import (
	"NyaLog/gin-blog-server/model"
	"testing"
)

func Test_ModifyPw(t *testing.T) {
	model.InitDb()
	model.ModifyPw("123456789")
}
