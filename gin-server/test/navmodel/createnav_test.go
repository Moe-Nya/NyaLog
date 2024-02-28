package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_CreateNav(t *testing.T) {
	model.InitDb()
	var navtitle = "归档"
	var navurl = "www.google.com"
	fmt.Println(model.CreateNav(navtitle, navurl))
}
