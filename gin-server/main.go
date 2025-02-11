package main

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/router"
	"NyaLog/gin-blog-server/service"
)

func init() {
	go service.StorageViews()
	go service.StorageLikes()
	go service.ClearUserLikes()
	go service.StorageCommentLikes()
	go service.SendCommentMail()
}

func main() {
	model.InitDb()
	router.IniRouter()
}
