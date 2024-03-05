package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestDeleArticle(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleteArticle(2))
}
