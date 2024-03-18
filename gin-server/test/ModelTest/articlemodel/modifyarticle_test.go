package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestModifyArticle(t *testing.T) {
	model.InitDb()
	var article = &model.Article{}
	article.Articleid = 2
	article.Articletitle = "新标题！"
	article.Cid = 3
	article.Aisummary = "321321"
	article.Text = "新文本！"
	fmt.Println(model.ModifyArticle(article))
}
