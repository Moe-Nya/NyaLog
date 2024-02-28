package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_ModifyFindme(t *testing.T) {
	model.InitDb()
	var findme = &model.FindMe{}
	findme.FindId = 0
	findme.Icon = "www.google.com"
	findme.Herf = "www.youtube.com"
	findme.Text = "油管"
	fmt.Println(model.ModifyFindme(findme))
}
