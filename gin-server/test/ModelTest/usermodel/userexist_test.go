package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestUserExist(t *testing.T) {
	model.InitDb()
	var r bool
	var e int
	r, e = model.UserExist()
	fmt.Println(r, e)
}
