package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSameCidArticle(t *testing.T) {
	model.InitDb()
	var article []model.Article
	article, _ = model.SameCidArt(1)
	fmt.Println(article)
}
