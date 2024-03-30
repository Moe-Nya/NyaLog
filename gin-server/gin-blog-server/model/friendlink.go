package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Friendlink struct {
	gorm.Model
	Friendlinkid      int    `gorm:"type:int(10);not null;primary_key" json:"friendlinkid" label:"友链id"`
	Friendname        string `gorm:"type:varchar(20);not null" json:"friendname" label:"站点名字"`
	Friendsite        string `gorm:"type:varchar(1000);not null" json:"friendsite" label:"站点地址"`
	Friendprofile     string `gorm:"type:varchar(1000)" json:"friendprofile" label:"站点图标"`
	Frienddescription string `gorm:"type:text" json:"frienddescription" label:"站点描述"`
}

// 新增友链
func CreateFriendLink(data *Friendlink) int {
	if data.Friendname == "" || data.Friendsite == "" {
		return errmsg.ERROR
	}
	var friendlink Friendlink
	friendlink.Friendname = data.Friendname
	friendlink.Friendsite = data.Friendsite
	friendlink.Friendprofile = data.Friendprofile
	friendlink.Frienddescription = data.Frienddescription
	var f Friendlink
	var count int64
	err := db.Find(&f).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	if count == 0 {
		friendlink.Friendlinkid = 0
	} else {
		var finalfriend Friendlink
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalfriend).Error
		if err != nil {
			return errmsg.ERROR
		}
		friendlink.Friendlinkid = finalfriend.Friendlinkid + 1
	}

	err = db.Create(&friendlink).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询友链列表
func SeleFriendLink(pageSize int, pageNum int) ([]Friendlink, int, int64) {
	var friendlinks []Friendlink
	err := db.Select("friendlink.friendlinkid, friendname, friendsite, friendprofile, created_at, frienddescription").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&friendlinks).Error
	if err != nil {
		return friendlinks, errmsg.ERROR, 0
	}
	var total int64
	db.Model(&friendlinks).Count(&total)
	return friendlinks, errmsg.SUCCESS, total
}

// 修改友链
func ModifyFriendLink(data *Friendlink) int {
	if data.Friendname == "" || data.Friendsite == "" {
		return errmsg.ERROR
	}
	var friendlink Friendlink
	var f Friendlink
	var friendlinkmaps = make(map[string]interface{})
	friendlinkmaps["friendname"] = data.Friendname
	friendlinkmaps["friendsite"] = data.Friendsite
	friendlinkmaps["friendprofile"] = data.Friendprofile
	friendlinkmaps["frienddescription"] = data.Frienddescription
	err := db.First(&f, "friendlinkid = ?", data.Friendlinkid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Model(&friendlink).Where("friendlinkid = ?", data.Friendlinkid).Updates(friendlinkmaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除友链
func DeleFriendLink(friendlinkid int) int {
	var friendlink Friendlink
	var f Friendlink
	err := db.First(&f, "friendlinkid = ?", friendlinkid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Where("friendlinkid = ?", friendlinkid).Unscoped().Delete(&friendlink).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
