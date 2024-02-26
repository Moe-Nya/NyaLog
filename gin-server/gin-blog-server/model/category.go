package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Cid          int    `gorm:"type:int(5);not null;primary key" json:"cid" label:"文章分类id"`
	Categorytext string `gorm:"type:varchar(50);not null" json:"categorytext" label:"分类内容"`
}
