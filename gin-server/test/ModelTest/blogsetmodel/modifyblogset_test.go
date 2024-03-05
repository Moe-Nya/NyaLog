package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestModifyBlogSet(t *testing.T) {
	model.InitDb()
	var blogset = &model.BlogSet{}
	blogset.Sitename = "猫猫头"
	blogset.Aiswitch = 0
	blogset.Aicategory = 1
	blogset.Commentswitch = 0
	blogset.Githuburl = "www.baidu.com"
	fmt.Println(model.ModifyBlogSet(blogset))
}
