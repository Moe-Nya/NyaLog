package model

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comid        string   `gorm:"type:varchar(20);not null;primary_key" json:"comid" label:"评论id"`
	Articleid    int64    `gorm:"type:bigint;not null" json:"articleid" label:"文章id"`
	Userid       string   `gorm:"type:varchar(20);not null" json:"userid" label:"用户id"`
	Profilephoto string   `gorm:"type:varchar(1000);not null" json:"profilephoto" label:"用户头像"`
	Usersite     string   `gorm:"type:varchar(1000)" json:"usersite" label:"用户主页"`
	Usermail     string   `gorm:"type:varchar(100)" json:"usermail" label:"用户邮箱"`
	Likes        string   `gorm:"type:varchar(1000);default:0" json:"likes" label:"点赞数"`
	Recomid      string   `gorm:"type:varchar(20)" json:"recomid" label:"回复评论id"`
	ReComInfo    *Comment `gorm:"type:json;default:null" json:"recominfo" label:"回复评论信息"`
	Commenttext  string   `gorm:"type:text;not null" json:"commenttext" label:"评论内容"`
}

// 新增评论
func CreateCom(comment *Comment) int {
	if comment.Userid == "" || comment.Commenttext == "" {
		return errmsg.ERROR
	}
	var article Article
	err := db.First(&article, "articleid = ?", comment.Articleid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	var count int64
	var c Comment
	err = db.Find(&c).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// 使id有一个自增效果，方便后续的寻找和修改\
	if count == 0 {
		comment.Comid = "0"
	} else {
		var finalcom Comment
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalcom).Error
		if err != nil {
			return errmsg.ERROR
		}
		comment.Comid = utils.BigNumAdd(finalcom.Comid)
	}
	err = db.Create(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询某文章的评论
func SeleCom(articleid int64, pageSize int, pageNum int) ([]*Comment, int) {
	var comments []*Comment
	var c Comment
	err := db.First(&c, "articleid = ?", articleid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return comments, errmsg.ERROR
	}
	err = db.Where("articleid = ?", articleid).Select("comment.comid, articleid, userid, profilephoto, likes, created_at, updated_at, usersite, recomid, commenttext").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&comments).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return comments, errmsg.SUCCESS
}

// 查询所有的评论
func SeleAllCom(pageSize int, pageNum int) ([]Comment, int, int64) {
	var commentList []Comment
	err := db.Select("comment.comid, articleid, userid, profilephoto, created_at, usersite, recomid, commenttext").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&commentList).Error
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	var total int64
	db.Model(&commentList).Count(&total)
	return commentList, errmsg.SUCCESS, total
}

// SelectOneCom 查询某篇文章的某个评论
func SelectOneCom(articleId int64, comId string) (*Comment, int) {
	var comment *Comment
	err := db.First(&comment, "articleid = ? and comid = ?", articleId, comId).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return comment, errmsg.SUCCESS
}

// 删除评论
func DeleteCom(comid string) int {
	var comment Comment
	var c Comment
	err := db.First(&c, "comid = ?", comid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Where("comid = ?", comid).Unscoped().Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除某篇文章的所有评论
func DeleteAllCom(articleid int64) int {
	var comment Comment
	var c Comment
	err := db.First(&c, "articleid = ?", articleid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Where("articleid = ?", articleid).Unscoped().Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// LikeCom 评论点赞
func LikeCom(comId string, likes string) int {
	var comment Comment
	err := db.First(&comment, "comid = ?", comId).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	likes = utils.BigDecimal(comment.Likes, likes)
	err = db.Model(&comment).Update("likes", likes).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
