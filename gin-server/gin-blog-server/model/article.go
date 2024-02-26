package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Articleid    int64      `gorm:"type:bigint;not null;primary key" json:"articleid" label:"文章id"`
	Articleimg   string     `gorm:"type:varchar(1000)" json:"articleimg" label:"文章头图"`
	Articletitle string     `gorm:"type:text;not null" json:"articletitle" label:"文章标题"`
	Articledate  time.Time  `gorm:"type:datetime;not null" json:"articledate" label:"文章最后更新时间"`
	Articlelikes string     `gorm:"type:varchar(100);not null;default:0" json:"articlelikes" label:"文章点赞数"`
	Articleviews string     `gorm:"type:varchar(100)not null;default:0" json:"articleviews" label:"文章浏览数"`
	Cid          []Category `gorm:"type:int;foreignKey:Cid" json:"cid" label:"文章分类id"`
	Aisummary    string     `gorm:"type:text" json:"aisummary" label:"AI文章摘要"`
	Text         string     `gorm:"type:text;not null" json:"text" label:"文章内容"`
}
