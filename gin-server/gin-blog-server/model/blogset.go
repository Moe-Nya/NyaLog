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
	Aiswitch       int       `gorm:"type:int(5);not null;default:0" json:"aiswitch" label:"AI摘要开关"`
	// 0是GPT 1是通义千问
	Aicategory    int `gorm:"type:int(5)" json:"aicategory" label:"使用谁的API"`
	Commentswitch int `gorm:"type:int(5);not null;default:0" json:"commentswitch" label:"是否开启评论"`
}

// 查询博客设置是否创建
func ValidateBlogSet() (int, int) {
	var count int64
	var b BlogSet
	err := db.Find(&b).Count(&count).Error
	if err != nil {
		return errmsg.ERROR, 0
	}
	if count < 1 {
		return errmsg.SUCCESS, 0
	}
	return errmsg.SUCCESS, 1
}

// 创建博客设置
func CreateBlogSet(data *BlogSet) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 读取博客设置信息
func SeleBlogSet() (BlogSet, int) {
	var blogset BlogSet
	e, ex := ValidateBlogSet()
	if e != errmsg.SUCCESS || ex == 0 {
		return blogset, errmsg.ERROR
	}
	err := db.Limit(1).Find(&blogset).Error
	if err != nil {
		return blogset, errmsg.ERROR
	}
	return blogset, errmsg.SUCCESS
}

// 读取博客设置信息(公开)
func QueryBlogSet() (BlogSet, int) {
	var blogset BlogSet
	e, ex := ValidateBlogSet()
	if e != errmsg.SUCCESS || ex == 0 {
		return blogset, errmsg.ERROR
	}
	err := db.Limit(1).Select("blog_set.sitename, sitecreatetime, sitebackground").Find(&blogset).Error
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
	}
	blogsetmap["commentswitch"] = data.Commentswitch
	var b BlogSet
	err := db.First(&b, "id = ?", 1).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Model(&blogset).Where("id = ?", 1).Updates(&blogsetmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
