package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_CreateNav(t *testing.T) {
	model.InitDb()
	var nav = &model.Navigation{}
	nav.Navtitle = "归档"
	nav.Navurl = "www.google.com"
	fmt.Println(model.CreateNav(nav))
}
