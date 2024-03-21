package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleCom(t *testing.T) {
	model.InitDb()
	var com []model.Comment
	com, _ = model.SeleCom(1, 1, 1)
	fmt.Println(com)
}
