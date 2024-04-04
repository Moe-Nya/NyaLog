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
	var blogset model.BlogSet
	blogset, err = model.SeleBlogSet()
	if err != errmsg.SUCCESS {
		return errmsg.SendCommentFailed
	} else {
		if blogset.Commentswitch == 0 {
			return errmsg.SendCommentFailed
		}
	}
	var com = &model.Comment{}
	com.Articleid = data.Articleid
	com.Commenttext = data.Comment
	com.Userid = userinfo["name"].(string)
	com.Usersite = userinfo["html_url"].(string)
	com.Profilephoto = userinfo["avatar_url"].(string)
	err = model.CreateCom(com)
	if err != errmsg.SUCCESS {
		return errmsg.SendCommentFailed
	}
	return errmsg.SUCCESS
}

// 读取文章评论所用结构体
type ArticleComments struct {
	Articleid int64  `json:"articleid"`
	Comid     string `json:"comid"`
	PageSize  int    `json:"pagesize"`
	PageNum   int    `json:"pagenum"`
}

// 读取评论列表(某篇文章)
func SeleCom(data *ArticleComments) ([]model.Comment, int) {
	var comment []model.Comment
	comment, err := model.SeleCom(data.Articleid, data.PageSize, data.PageNum)
	if err != errmsg.SUCCESS {
		return comment, errmsg.SelectCommentFailed
	}
	return comment, errmsg.SUCCESS
}

// 读取评论列表(所有文章)
func SeleAllCom(data *ArticleComments) ([]model.Comment, int, int64) {
	var comment []model.Comment
	comment, err, total := model.SeleAllCom(data.PageSize, data.PageNum)
	if err != errmsg.SUCCESS {
		return comment, errmsg.SelectCommentFailed, 0
	}
	return comment, errmsg.SUCCESS, total
}

// 删除评论(单个评论)
func DeleteCom(data *ArticleComments) int {
	err := model.DeleteCom(data.Comid)
	if err != errmsg.SUCCESS {
		return errmsg.DeleteCommentFailed
	}
	return errmsg.SUCCESS
}

// 删除评论(某篇文章的所有评论)
func DeleteAllCom(articleid int64) int {
	err := model.DeleteAllCom(articleid)
	if err != errmsg.SUCCESS {
		return errmsg.DeleteCommentFailed
	}
	return errmsg.SUCCESS
}
