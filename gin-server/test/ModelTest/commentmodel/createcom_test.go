package main

import (
	"NyaLog/gin-blog-server/model"
	"fmt"
	"testing"
)

func TestCreateCom(t *testing.T) {
	model.InitDb()
	var com = &model.Comment{}
	com.Articleid = 1
	com.Userid = "moenya"
	com.Profilephoto = "www.google.com"
	com.Usersite = "www.youtube.com"
	com.Commenttext = "测试回复2"
	com.Recomid = "0"
	fmt.Println(model.CreateCom(com))
}
