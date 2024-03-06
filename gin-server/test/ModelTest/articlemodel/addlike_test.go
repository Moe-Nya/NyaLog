package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestAddLike(t *testing.T) {
	model.InitDb()
	fmt.Println(model.AddLike(0))
}
