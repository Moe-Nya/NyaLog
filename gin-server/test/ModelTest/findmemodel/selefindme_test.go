package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Findme(t *testing.T) {
	model.InitDb()
	var findmes []model.FindMe
	findmes, _ = model.SeleFindme()
	fmt.Println(findmes)
}
