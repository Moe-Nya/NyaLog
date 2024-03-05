package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BlogSet struct {
	gorm.Model
	Sitename       string    `gorm:"type:varchar(50)" json:"sitename" label:"站点名称"`
	Sitecreatetime time.Time `gorm:"type:datetime" json:"sitecreatetime" label:"创建时间"`
	Sitebackground string    `gorm:"type:varchar(1000)" json:"sitebackground" label:"博客头图"`
	Aiswitch       int       `gorm:"type:int(5);not null;default:0" json:"aiswitch" label:"AI摘要开关"`
	// 0是GPT 1是通义千问
	Aicategory    int    `gorm:"type:int(5)" json:"aicategory" label:"使用谁的API"`
	Aiurl         string `gorm:"type:varchar(1000)" json:"aiurl" label:"API地址"`
	Commentswitch int    `gorm:"type:int(5);not null" json:"commentswitch" label:"是否开启评论"`
	Githuburl     string `gorm:"type:varchar(1000)" json:"githuburl" label:"githubAPI"`
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
	blogsetmap["aiswitch"] = data.Aiswitch
	if blogsetmap["aiswitch"] == 1 {
		blogsetmap["aicategory"] = data.Aicategory
		blogsetmap["aiurl"] = data.Aiurl
	}
	blogsetmap["commentswitch"] = data.Commentswitch
	if blogsetmap["commentswitch"] == 1 {
		blogsetmap["githuburl"] = data.Githuburl
	}
	err := db.Model(&blogset).Where("id = ?", 1).Updates(&blogsetmap).Error
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
