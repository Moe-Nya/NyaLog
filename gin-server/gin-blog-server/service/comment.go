package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// Comment 评论所需要的结构体
type Comment struct {
	Articleid   int64  `json:"articleid"`
	Comment     string `json:"comment"`
	ReCommentId string `json:"replayid"`
}

// 添加评论
func NewComment(data *Comment, token string) int {
	userinfo, err := middleware.GetUserInfo(token)
	if err != errmsg.SUCCESS {
		return err
	}
	var blogset model.BlogSet
	blogset, err = model.SeleBlogSet()
	if err != errmsg.SUCCESS || blogset.Commentswitch == 0 {
		return errmsg.SendCommentFailed
	}
	var com = &model.Comment{}
	if userinfo["name"] == nil {
		if userinfo["login"] == nil || userinfo["avatar_url"] == nil || userinfo["html_url"] == nil {
			return errmsg.SendCommentFailed
		} else {
			com.Userid = userinfo["login"].(string)
		}
	} else {
		com.Userid = userinfo["name"].(string)
	}
	com.Articleid = data.Articleid
	com.Commenttext = data.Comment
	com.Recomid = data.ReCommentId
	com.Usersite = userinfo["html_url"].(string)
	com.Profilephoto = userinfo["avatar_url"].(string)
	err = model.CreateCom(com)
	if err != errmsg.SUCCESS {
		return errmsg.SendCommentFailed
	}
	return errmsg.SUCCESS
}

// 评论用户信息结构体
type CommentUser struct {
	Username    string `json:"username"`
	Userprofile string `json:"userprofile"`
	Usersite    string `json:"usersite"`
}

// 查询评论用户信息
func QueryCommentUserInfo(token string) (CommentUser, int) {
	var comuser CommentUser
	userinfo, err := middleware.GetUserInfo(token)
	if err != errmsg.SUCCESS {
		return comuser, errmsg.GetQueryFailed
	}
	if userinfo["name"] == nil {
		if userinfo["login"] == nil || userinfo["avatar_url"] == nil || userinfo["html_url"] == nil {
			return comuser, errmsg.GetQueryFailed
		} else {
			comuser.Username = userinfo["login"].(string)
		}
	} else {
		comuser.Username = userinfo["name"].(string)
	}
	comuser.Userprofile = userinfo["avatar_url"].(string)
	comuser.Usersite = userinfo["html_url"].(string)
	return comuser, errmsg.SUCCESS
}

// 读取文章评论所用结构体
type ArticleComments struct {
	Articleid int64  `json:"articleid"`
	Comid     string `json:"comid"`
	PageSize  int    `json:"pagesize"`
	PageNum   int    `json:"pagenum"`
}

// 读取评论列表(某篇文章)
func SeleCom(data *ArticleComments) ([]*model.Comment, int) {
	var comment []*model.Comment
	comment, err := model.SeleCom(data.Articleid, data.PageSize, data.PageNum)
	for _, com := range comment {
		if middleware.CommentLikeInfoMap[com.Comid] != nil && com.Likes != "" {
			com.Likes = utils.BigDecimal(com.Likes, middleware.CommentLikeInfoMap[com.Comid].Likes)
		}
		if com.Recomid != "" {
			var recom *model.Comment
			recom, err = model.SelectOneCom(com.Articleid, com.Recomid)
			if err != errmsg.SUCCESS {
				continue
			}
			com.ReComInfo = recom
		}
	}
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

type CommentLikeInfo struct {
	ComId string `json:"comid"`
}

// LikeCom 评论点赞
func LikeCom(com CommentLikeInfo, ip string) int {
	err := AddCommentLike(com.ComId, ip)
	if err != errmsg.SUCCESS {
		return errmsg.LikeCommentFailed
	}
	return errmsg.SUCCESS
}
