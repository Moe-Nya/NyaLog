package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestDeleCat(t *testing.T) {
	model.InitDb()
	fmt.Println(model.DeleCat(1))
}
