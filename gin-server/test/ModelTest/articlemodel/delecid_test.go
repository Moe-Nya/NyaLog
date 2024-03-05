package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestDeleCid(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleCid(1))
}
