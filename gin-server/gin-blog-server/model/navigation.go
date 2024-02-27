package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Navigation struct {
	gorm.Model
	NavId    int    `gorm:"type:int(10);not null" json:"navid" label:"导航标签id"`
	Navtitle string `gorm:"type:varchar(20);not null" json:"navtitle" label:"导航栏标题"`
	Navurl   string `gorm:"type:varchar(1000)" json:"navurl" label:"导航链接"`
}

// 增加导航标签
func CreateNav(navtitle string, navurl string) int {
	var nav Navigation
	nav.Navtitle = navtitle
	nav.Navurl = navurl
	err := db.Create(&nav).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取导航标签列表
func SeleNav() (Navigation, int) {
	var nav Navigation
	err := db.Limit(1).Find(&nav).Error
	if err != nil {
		return nav, errmsg.ERROR
	}
	return nav, errmsg.SUCCESS
}

// 修改导航标签

// 删除导航标签
