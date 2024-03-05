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
	blogset.Aiurl = "opai.com"
	blogset.Commentswitch = 1
	blogset.Githuburl = "www.github.com"
	fmt.Println(model.CreateBlogSet(blogset))
}
