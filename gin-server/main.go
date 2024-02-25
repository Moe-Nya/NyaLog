package main

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/router"
)

func main() {
	model.InitDb()
	router.IniRouter()
}
