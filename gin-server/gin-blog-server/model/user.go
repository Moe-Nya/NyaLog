package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid          string `gorm:"type:varchar(20);not null;primary key" json:"uid" label:"用户UID"`
	Username     string `gorm:"type:varchar(20);not null" json:"username" label:"用户名"`
	Password     string `gorm:"type:varchar(100);not null" json:"password" label:"用户密码"`
	Profilephoto string `gorm:"type:varchar(1000)" json:"profilephoto" label:"用户头像"`
	Email        string `gorm:"type:varchar(50);not null" json:"email" label:"用户邮箱"`
	Slogan       string `gorm:"type:varchar(50)" json:"slogan" label:"slogan"`
	Salt         string `gorm:"type:varchar(50);not null" json:"salt" label:"salt"`
	Secret       string `gorm:"type:varchar(300);not null" json:"secret" label:"secret"`
	Lastip       string `gorm:"type:varchar(20)" json:"lastip" label:"lastip"`
}

// 新增用户
func NewUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 提取用户信息

// 更新用户信息
