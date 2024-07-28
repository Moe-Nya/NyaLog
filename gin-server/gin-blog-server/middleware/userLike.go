package middleware

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"sync"
)

// 处理用户like池结构体
type UserLike struct {
	Likes    string
	UserLike map[string]bool
}

// 以articleid为索引来记录like池
var UserLikeMap map[int64]*UserLike

// 初始化like池
func init() {
	UserLikeMap = make(map[int64]*UserLike)
}

// 增加文章的临时like量
var mutexUserLike sync.Mutex

func AddLike(articleid int64, ip string) int {
	mutexUserLike.Lock()
	defer mutexUserLike.Unlock()
	// 确保 UserLikeMap[articleid] 不为 nil
	if UserLikeMap[articleid] == nil {
		UserLikeMap[articleid] = &UserLike{
			Likes:    "",
			UserLike: make(map[string]bool),
		}
	}
	if !UserLikeMap[articleid].UserLike[ip] {
		UserLikeMap[articleid].UserLike[ip] = true
		if UserLikeMap[articleid].Likes == "" {
			UserLikeMap[articleid].Likes = "1"
		} else {
			UserLikeMap[articleid].Likes = utils.BigNumAdd(UserLikeMap[articleid].Likes)
		}
		return errmsg.SUCCESS
	} else {
		return errmsg.ERROR
	}
}

// 清空map并且存库like
func ClearLikes() int {
	mutexUserLike.Lock()
	defer mutexUserLike.Unlock()
	for index, v := range UserLikeMap {
		err := model.AddLike(index, v.Likes)
		if err != errmsg.SUCCESS {
			return errmsg.LikesBufferStorageFailed
		}
		UserLikeMap[index].Likes = ""
	}
	return errmsg.SUCCESS
}

// 清空用户点赞记录
func ClearUserLike() {
	mutexUserLike.Lock()
	defer mutexUserLike.Unlock()
	for index1, v := range UserLikeMap {
		for index2 := range v.UserLike {
			delete(UserLikeMap[index1].UserLike, index2)
		}
	}
}
