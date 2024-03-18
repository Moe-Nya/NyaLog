package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
	"time"
)

func TestModifyBlogSet(t *testing.T) {
	model.InitDb()
	var blogset = &model.BlogSet{}
	blogset.Sitecreatetime = time.Now()
	blogset.Sitename = "猫猫头"
	blogset.Aiswitch = 0
	blogset.Aicategory = 1
	blogset.Commentswitch = 0
	fmt.Println(model.ModifyBlogSet(blogset))
}
