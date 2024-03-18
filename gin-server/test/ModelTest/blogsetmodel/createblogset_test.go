package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
	"time"
)

func TestCreateBlogSet(t *testing.T) {
	model.InitDb()
	var blogset = &model.BlogSet{}
	blogset.Sitename = "MMT"
	blogset.Sitecreatetime = time.Now()
	blogset.Sitebackground = "www.google.com"
	blogset.Aiswitch = 1
	blogset.Aicategory = 0
	blogset.Commentswitch = 1
	fmt.Println(model.CreateBlogSet(blogset))
}
