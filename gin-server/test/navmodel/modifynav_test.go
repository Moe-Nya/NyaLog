package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_ModifyNav(t *testing.T) {
	model.InitDb()
	var navid = 2
	var navtitle = "友链"
	var navurl = "www.google.com"
	fmt.Println(model.ModifyNav(navid, navtitle, navurl))
}
