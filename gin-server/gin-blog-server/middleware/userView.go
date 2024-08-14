package middleware

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"sync"
)

// 用户浏览量和IP访问情况
type UserView struct {
	Views    string
	UserView map[string]bool
}

// 临时存储每篇文章的浏览量
var ArticleViewInfo map[int64]*UserView

var mutexUserView sync.Mutex

// 对该存储映射进行初始化
func init() {
	ArticleViewInfo = make(map[int64]*UserView)
}

// 添加文章id
func AddArticleId(articleid int64, ip string) bool {
	mutexUserView.Lock()
	defer mutexUserView.Unlock()
	if ArticleViewInfo[articleid] == nil {
		ArticleViewInfo[articleid] = &UserView{
			Views:    "",
			UserView: make(map[string]bool),
		}
	}
	if !ArticleViewInfo[articleid].UserView[ip] {
		ArticleViewInfo[articleid].UserView[ip] = true
		if ArticleViewInfo[articleid].Views == "" {
			ArticleViewInfo[articleid].Views = "1"
		} else {
			ArticleViewInfo[articleid].Views = utils.BigNumAdd(ArticleViewInfo[articleid].Views)
		}
		return true
	} else {
		return false
	}
}

// 存库并且清空浏览量
func ClearView() int {
	mutexUserView.Lock()
	defer mutexUserView.Unlock()
	for index, v := range ArticleViewInfo {
		err := model.AddViews(index, v.Views)
		if err != errmsg.SUCCESS {
			return errmsg.ViewsBufferStorageFailed
		}
		delete(ArticleViewInfo, index)
	}
	return errmsg.SUCCESS
}
