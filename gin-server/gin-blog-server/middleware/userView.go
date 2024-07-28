package middleware

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"sync"
)

// 临时存储每篇文章的浏览量
var ArticleViewInfo map[int64]string

var mutexUserView sync.Mutex

// 对该存储映射进行初始化
func init() {
	ArticleViewInfo = make(map[int64]string)
}

// 添加文章id
func AddArticleId(articleid int64) bool {
	mutexUserView.Lock()
	defer mutexUserView.Unlock()
	if ArticleViewInfo[articleid] == "" {
		ArticleViewInfo[articleid] = "1"
		return false
	}
	return true
}

// 增加浏览量
func AddView(articleid int64) {
	mutexUserView.Lock()
	defer mutexUserView.Unlock()
	ArticleViewInfo[articleid] = utils.BigNumAdd(ArticleViewInfo[articleid])
}

// 存库并且清空浏览量
func ClearView() int {
	mutexUserView.Lock()
	defer mutexUserView.Unlock()
	for index, v := range ArticleViewInfo {
		err := model.AddViews(index, v)
		if err != errmsg.SUCCESS {
			return errmsg.ViewsBufferStorageFailed
		}
		delete(ArticleViewInfo, index)
	}
	return errmsg.SUCCESS
}
