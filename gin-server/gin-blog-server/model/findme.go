package model

import "gorm.io/gorm"

type FindMe struct {
	gorm.Model
	FindId int    `gorm:"type:int(10);not null" json:"findid" label:"找到我ID"`
	Icon   string `gorm:"type:varchar(1000);not null" json:"icon" label:"图标"`
	Herf   string `gorm:"type:varchar(1000)not null" json:"herf" label:"链接"`
	Text   string `gorm:"type:varchar(50)" json:"text" label:"文本"`
}

// 新增Findme

// 查询所有Findme

// 修改Findme

// 删除Findme
