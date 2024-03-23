package model

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"fmt"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comid        string `gorm:"type:varchar(20);not null;primary_key" json:"comid" label:"评论id"`
	Articleid    int64  `gorm:"type:bigint;not null" json:"articleid" label:"文章id"`
	Userid       string `gorm:"type:varchar(20);not null" json:"userid" label:"用户id"`
	Profilephoto string `gorm:"type:varchar(1000);not null" json:"profilephoto" label:"用户头像"`
	Usersite     string `gorm:"type:varchar(1000)" json:"usersite" label:"用户主页"`
	Recomid      string `gorm:"type:varchar(20)" json:"recomid" label:"回复评论id"`
	Commenttext  string `gorm:"type:text;not null" json:"commenttext" label:"评论内容"`
}

// 新增评论
func CreateCom(comment *Comment) int {
	var article Article
	e := db.First(&article, "articleid = ?", comment.Articleid)
	if e.Error != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	var count int64
	var c Comment
	err := db.Find(&c).Count(&count).Error
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
func SeleCom(articleid int64, pageSize int, pageNum int) ([]Comment, int) {
	var comments []Comment
	var c Comment
	err := db.First(&c, "articleid = ?", articleid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return comments, errmsg.ERROR
	}
	err = db.Where("articleid = ?", articleid).Select("comment.comid, userid, profilephoto, created_at, updated_at, usersite, recomid, commenttext").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&comments).Error
	fmt.Println(comments)
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
