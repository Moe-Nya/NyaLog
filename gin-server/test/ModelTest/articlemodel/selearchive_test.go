package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestSeleArchive(t *testing.T) {
	model.InitDb()
	var archive []model.Archive
	archive, err := model.SeleArchive()
	fmt.Println(archive, err)
}
