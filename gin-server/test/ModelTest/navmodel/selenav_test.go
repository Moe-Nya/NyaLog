package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_SeleNav(t *testing.T) {
	model.InitDb()
	var nav []model.Navigation
	nav, _ = model.SeleNav()
	for _, i := range nav {
		fmt.Println(i.NavId)
	}
}
