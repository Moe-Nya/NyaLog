package model

import "gorm.io/gorm"

type Navigation struct {
	gorm.Model
	Navtitle string `gorm:"type:varchar(20);not null" json:"navtitle" label:"导航栏标题"`
	Navurl   string `gorm:"type:varchar(1000)" json:"navurl" label:"导航链接"`
}
