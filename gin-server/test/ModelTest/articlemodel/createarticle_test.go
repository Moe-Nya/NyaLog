package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	model.InitDb()
	var article = &model.Article{}
	article.Articletitle = "喵喵喵"
	article.Articleimg = "www.google.com"
	article.Articlelikes = "123"
	article.Articleviews = "456"
	article.Cid = 1
	article.Aisummary = "not fund"
	article.Text = "12312321312312312444"
	article.Shorttext = "1234"
	fmt.Println(model.CreateArticle(article))
}
