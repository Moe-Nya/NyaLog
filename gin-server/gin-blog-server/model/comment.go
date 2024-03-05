package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Comment struct {
	Article Article `gorm:"foreignkey:Article"`
	gorm.Model
	Comid        string `gorm:"type:varchar(20);not null;primary key" json:"comid" label:"评论id"`
	Articleid    int64  `gorm:"type:bigint;not null;" json:"articleid" label:"文章id"`
	Userid       string `gorm:"type:varchar(20);not null" json:"userid" label:"用户id"`
	Profilephoto string `gorm:"type:varchar(1000);not null" json:"profilephoto" label:"用户头像"`
	Usersite     string `gorm:"type:varchar(1000)" json:"usersite" label:"用户主页"`
	Recomid      string `gorm:"type:varchar(20)" json:"recomid" label:"回复评论id"`
	Comment      string `gorm:"type:text;not null" json:"comment" label:"评论内容"`
}

// 新增评论
func CreateCom(comment *Comment) int {
	err := db.Where("articleid = ?", comment.Articleid).Create(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询某文章的评论
func SeleCom(articleid int64) ([]Comment, int) {
	var comment []Comment
	err := db.Where("articleid = ?", articleid).Find(&comment).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return comment, errmsg.SUCCESS
}

// 删除评论
func DeleteCom(Comid string) int {
	var comment Comment
	err := db.Where("Comid = ?", Comid).Unscoped().Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
