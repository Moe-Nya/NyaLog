package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestDeleAllCom(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleteAllCom(1))
}
