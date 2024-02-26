package model

import "gorm.io/gorm"

type FindMe struct {
	gorm.Model
	Icon string `gorm:"type:varchar(1000);not null" json:"icon" label:"图标"`
	Herf string `gorm:"type:varchar(1000)not null" json:"herf" label:"链接"`
	Text string `gorm:"type:varchar(50)" json:"text" label:"文本"`
}
