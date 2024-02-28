package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func Test_Delefindme(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleFindme(0))
}
