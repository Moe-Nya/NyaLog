package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestDeleteCom(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleteCom("0"))
}
