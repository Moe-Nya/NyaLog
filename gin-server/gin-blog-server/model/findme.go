package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type FindMe struct {
	gorm.Model
	FindId int    `gorm:"type:int(10);not null;primary_key" json:"findid" label:"找到我ID"`
	Icon   string `gorm:"type:varchar(1000);not null" json:"icon" label:"图标"`
	Herf   string `gorm:"type:varchar(1000)not null" json:"herf" label:"链接"`
	Text   string `gorm:"type:varchar(50)" json:"text" label:"文本"`
}

// 新增Findme
func CreateFindme(data *FindMe) int {
	var findme FindMe
	findme.Icon = data.Icon
	findme.Herf = data.Herf
	findme.Text = data.Text
	var count int64
	var f FindMe
	err := db.Find(&f).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// id自增效果计算
	if count == 0 {
		findme.FindId = 0
	} else {
		var finalfind FindMe
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalfind).Error
		if err != nil {
			return errmsg.ERROR
		}
		findme.FindId = finalfind.FindId + 1
	}
	err = db.Create(&findme).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询所有Findme
func SeleFindme() ([]FindMe, int) {
	var findme []FindMe
	err := db.Find(&findme).Error
	if err != nil {
		return findme, errmsg.ERROR
	}
	return findme, errmsg.SUCCESS
}

// 修改Findme
func ModifyFindme(data *FindMe) int {
	var findme FindMe
	var findmap = make(map[string]interface{})
	findmap["icon"] = data.Icon
	findmap["herf"] = data.Herf
	findmap["text"] = data.Text
	err := db.Model(&findme).Where("find_id = ?", data.FindId).Updates(findmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除Findme
func DeleFindme(findid int) int {
	var findme FindMe
	err := db.Where("find_id = ?", findid).Unscoped().Delete(&findme).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
