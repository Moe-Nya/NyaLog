package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 评论所需要的结构体
type Comment struct {
	Articleid int64  `json:"articleid"`
	Comment   string `json:"comment"`
}

// 添加评论
func NewComment(data *Comment, token string) int {
	userinfo, err := middleware.GetUserInfo(token)
	if err != errmsg.SUCCESS {
		return err
	}
	var com = &model.Comment{}
	com.Articleid = data.Articleid
	com.Comment = data.Comment
	com.Userid = userinfo["name"].(string)
	com.Usersite = userinfo["html_url"].(string)
	com.Profilephoto = userinfo["avatar_url"].(string)
	err = model.CreateCom(com)
	if err != errmsg.SUCCESS {
		return errmsg.SendCommentFailed
	}
	return errmsg.SUCCESS
}

// 读取评论列表(某篇文章)
