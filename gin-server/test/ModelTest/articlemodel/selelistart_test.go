package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleListArt(t *testing.T) {
	model.InitDb()
	var article = []model.Article{}
	var total int64
	article, _, total = model.SeleListArticle(1, 2)
	fmt.Println(article[0].Articleid, total)
}
