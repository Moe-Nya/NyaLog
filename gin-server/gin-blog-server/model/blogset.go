package model

import (
	"time"

	"gorm.io/gorm"
)

type BlogSet struct {
	gorm.Model
	Sitename       string    `gorm:"type:varchar(50);not null" json:"sitename" label:"站点名称"`
	Sitecreatetime time.Time `gorm:"type:datetime;not null" json:"sitecreatetime" label:"创建时间"`
}
