package service

import (
	"NyaLog/gin-blog-server/middleware"
	"time"
)

// 存储like
func StorageLikes() {
	var counter int = 0
	for {
		time.Sleep(time.Hour)
		counter++
		if counter >= 2 {
			counter = 0
			middleware.ClearLikes()
		}
	}
}

// 清空用户点赞记录
func ClearUserLikes() {
	var counter int = 0
	for {
		time.Sleep(time.Hour)
		counter++
		if counter >= 24 {
			counter = 0
			middleware.ClearUserLike()
		}
	}
}
