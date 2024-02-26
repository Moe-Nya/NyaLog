package model

import "gorm.io/gorm"

type BlogBackSet struct {
	gorm.Model
	Sitebackground string `gorm:"type:varchar(1000)" json:"sitebackground" label:"博客头图"`
}
