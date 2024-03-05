package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestModifyCat(t *testing.T) {
	model.InitDb()
	var cat1 []model.Category
	cat1, _ = model.SeleCat()
	var cat2 = &model.Category{}
	for _, i := range cat1 {
		if i.Cid == 1 {
			*cat2 = i
		}
	}
	cat2.Categorytext = "life"
	fmt.Println(model.ModifyCat(cat2))
}
