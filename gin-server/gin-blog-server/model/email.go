package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type EmailServer struct {
	gorm.Model
	Stmphost      string `gorm:"type:varchar(100);not null" json:"stmphost" label:"邮件服务器地址"`
	Stmpport      int    `gorm:"type:int(10);not null" json:"stmpport" label:"邮件服务器端口"`
	Emailusername string `gorm:"type:varchar(50);not null" json:"emailusername" label:"邮件服务器用户名"`
	Emailpassword string `gorm:"type:varchar(50)" json:"emailpassword" label:"邮件服务器密码"`
}

// 创建邮件服务器信息
func CreateEmai(data *EmailServer) int {
	var count int64
	var e EmailServer
	// 先查询数据库中邮件服务器信息的条数（只能有一条）
	err := db.Find(&e).Count(&count).Error
	if err != nil || count > 0 {
		return errmsg.ERROR
	}
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 读取邮件服务器信息
func SeleEmail() (EmailServer, int) {
	var email EmailServer
	err := db.Limit(1).Find(&email).Error
	if err != nil {
		return email, errmsg.ERROR
	}
	return email, errmsg.SUCCESS
}

// 修改邮件服务器信息
func ModifyEmail(data *EmailServer) int {
	var email EmailServer
	var emailmap = make(map[string]interface{})
	emailmap["stmphost"] = data.Stmphost
	emailmap["stmpport"] = data.Stmpport
	emailmap["emailusername"] = data.Emailusername
	emailmap["emailpassword"] = data.Emailpassword
	err := db.Model(&email).Where("id = 1").Updates(emailmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
