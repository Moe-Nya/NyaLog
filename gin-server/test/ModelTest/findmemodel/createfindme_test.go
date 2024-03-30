package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_CreateFindme(t *testing.T) {
	model.InitDb()
	var findme = &model.FindMe{}
	findme.Icon = "www.google.com"
	findme.Href = "www.bilibili.com"
	findme.Text = "哔哩哔哩"
	fmt.Println(model.CreateFindme(findme))
}
