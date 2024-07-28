package service

import (
	"NyaLog/gin-blog-server/middleware"
	"time"
)

// 记录浏览新增
func AddView(articleid int64) {
	b := middleware.AddArticleId(articleid)
	if b {
		middleware.AddView(articleid)
	}
}

// 存库
func StorageViews() {
	var counter int
	for {
		time.Sleep(time.Hour)
		counter++
		if counter >= 2 {
			counter = 0
			middleware.ClearView()
		}
	}
}
