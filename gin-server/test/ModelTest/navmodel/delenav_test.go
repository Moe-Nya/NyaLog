package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_DeleNav(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleNav(3))
}
