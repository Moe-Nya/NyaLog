package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_ModifyNav(t *testing.T) {
	model.InitDb()
	var nav *model.Navigation
	nav.NavId = 2
	nav.Navtitle = "友链"
	nav.Navurl = "www.google.com"
	fmt.Println(model.ModifyNav(nav))
}
