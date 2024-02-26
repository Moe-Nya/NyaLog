package model

import "gorm.io/gorm"

type EmailServer struct {
	gorm.Model
	Stmphost      string `gorm:"type:varchar(100);not null" json:"stmphost" label:"邮件服务器地址"`
	Stmpport      int    `gorm:"type:int(10);not null" json:"stmpport" label:"邮件服务器端口"`
	Emailusername string `gorm:"type:varchar(50);not null" json:"emailusername" label:"邮件服务器用户名"`
	Emailpassword string `gorm:"type:varchar(50)" json:"emailpassword" label:"邮件服务器密码"`
}
