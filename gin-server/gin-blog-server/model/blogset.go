package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"time"

	"gorm.io/gorm"
)

type BlogSet struct {
	gorm.Model
	Sitename       string    `gorm:"type:varchar(50)" json:"sitename" label:"站点名称"`
	Sitecreatetime time.Time `gorm:"type:datetime" json:"sitecreatetime" label:"创建时间"`
	Sitebackground string    `gorm:"type:varchar(1000)" json:"sitebackground" label:"博客头图"`
}

// 创建博客设置
func CreateBlogSet(data *BlogSet) int {
	var count int64
	var b BlogSet
	err := db.Find(&b).Count(&count).Error
	if err != nil || count > 0 {
		return errmsg.ERROR
	}
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 读取博客设置信息
func SeleBlogSet() (BlogSet, int) {
	var blogset BlogSet
	err := db.Limit(1).Find(&blogset).Error
	if err != nil {
		return blogset, errmsg.ERROR
	}
	return blogset, errmsg.SUCCESS
}

// 修改博客设置信息
func ModifyBlogSet(data *BlogSet) int {
	var blogset BlogSet
	var blogsetmap = make(map[string]interface{})
	blogsetmap["sitename"] = data.Sitename
	blogsetmap["sitecreatetime"] = data.Sitecreatetime
	blogsetmap["sitebackground"] = data.Sitebackground
	err := db.Model(&blogset).Where("id = ?", data.ID).Updates(blogsetmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
