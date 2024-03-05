package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleCat(t *testing.T) {
	model.InitDb()
	fmt.Println(model.SeleCat())
}
