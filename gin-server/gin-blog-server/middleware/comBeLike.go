package middleware

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"sync"
)

type CommentLikeInfo struct {
	Likes       string
	CommentLike map[string]bool
}

var CommentLikeInfoMap map[string]*CommentLikeInfo

var mutexCommentLike sync.Mutex

func init() {
	CommentLikeInfoMap = make(map[string]*CommentLikeInfo)
}

func AddCommentLike(comId string, ip string) int {
	mutexCommentLike.Lock()
	defer mutexCommentLike.Unlock()
	if CommentLikeInfoMap[comId] == nil {
		var commentLike = make(map[string]bool)
		commentLike[ip] = true
		CommentLikeInfoMap[comId] = &CommentLikeInfo{
			Likes:       "1",
			CommentLike: commentLike,
		}
		return errmsg.SUCCESS
	}
	if !CommentLikeInfoMap[comId].CommentLike[ip] {
		CommentLikeInfoMap[comId].CommentLike[ip] = true
		if CommentLikeInfoMap[comId].Likes == "" {
			CommentLikeInfoMap[comId].Likes = "1"
		} else {
			CommentLikeInfoMap[comId].Likes = utils.BigDecimal(CommentLikeInfoMap[comId].Likes, "1")
		}
		return errmsg.SUCCESS
	}
	return errmsg.LikeCommentFailed
}

func ClearCommentLike() int {
	mutexCommentLike.Lock()
	defer mutexCommentLike.Unlock()
	for index, v := range CommentLikeInfoMap {
		err := model.LikeCom(index, v.Likes)
		if err != errmsg.SUCCESS {
			return errmsg.LikeBufferStorageFailed
		}
		delete(CommentLikeInfoMap, index)
	}
	return errmsg.SUCCESS
}
