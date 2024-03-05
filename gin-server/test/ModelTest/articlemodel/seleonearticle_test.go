package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleOneArticle(t *testing.T) {
	model.InitDb()
	var article model.Article
	article, _ = model.SeleOneArticle(0)
	fmt.Println(article)
}
