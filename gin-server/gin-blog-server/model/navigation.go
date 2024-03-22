package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Navigation struct {
	gorm.Model
	NavId    int    `gorm:"type:int(10);not null;primary_key" json:"navid" label:"导航标签id"`
	Navtitle string `gorm:"type:varchar(20);not null" json:"navtitle" label:"导航栏标题"`
	Navurl   string `gorm:"type:varchar(1000)" json:"navurl" label:"导航链接"`
}

// 增加导航标签
func CreateNav(data *Navigation) int {
	var nav Navigation
	nav.Navtitle = data.Navtitle
	nav.Navurl = data.Navurl
	var count int64
	// 单独创建一个结构体，因为db查询后会把查询结果放到n中，如果这里使用db.Find(&nav)，就会使
	// 最后一条数据被赋值给nav，ID相同导致创建失败
	var n Navigation
	err := db.Find(&n).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// 使navid有一个自增效果，方便后续的寻找和修改
	if count == 0 {
		nav.NavId = 0
	} else {
		var finalnav Navigation
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalnav).Error
		if err != nil {
			return errmsg.ERROR
		}
		nav.NavId = finalnav.NavId + 1
	}
	err = db.Create(&nav).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询所有导航标签
func SeleNav() ([]Navigation, int) {
	var nav []Navigation
	err := db.Find(&nav).Error
	if err != nil {
		return nav, errmsg.ERROR
	}
	return nav, errmsg.SUCCESS
}

// 修改导航标签
func ModifyNav(data *Navigation) int {
	var nav Navigation
	var navmap = make(map[string]interface{})
	navmap["navtitle"] = data.Navtitle
	navmap["navurl"] = data.Navurl
	var n Navigation
	err := db.First(&n, "nav_id = ?", data.NavId).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Model(&nav).Where("nav_id = ?", data.NavId).Updates(navmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除导航标签
func DeleNav(navid int) int {
	var nav Navigation
	var n Navigation
	err := db.First(&n, "nav_id = ?", navid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Where("nav_id = ?", navid).Unscoped().Delete(&nav).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
