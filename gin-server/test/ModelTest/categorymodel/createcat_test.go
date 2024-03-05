package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestCreateCat(t *testing.T) {
	model.InitDb()
	var cat = &model.Category{}
	cat.Categorytext = "生活"
	fmt.Println(model.CreateCat(cat))
}
