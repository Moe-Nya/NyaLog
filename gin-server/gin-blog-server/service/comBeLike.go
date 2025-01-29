package service

import (
	"NyaLog/gin-blog-server/middleware"
	"time"
)

// AddCommentLike 增加评论点赞
func AddCommentLike(commentId string, ip string) int {
	err := middleware.AddCommentLike(commentId, ip)
	return err
}

// StorageCommentLikes 存储评论点赞并清除缓存
func StorageCommentLikes() {
	var counter int = 0
	for {
		time.Sleep(time.Hour)
		counter++
		if counter >= 6 {
			counter = 0
			middleware.ClearCommentLike()
		}
	}
}
