package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleBlogSet(t *testing.T) {
	model.InitDb()
	fmt.Println(model.SeleBlogSet())
}
