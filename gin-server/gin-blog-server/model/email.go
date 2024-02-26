package model

import "gorm.io/gorm"

type EmailServer struct {
	gorm.Model
	Stmphost      string `gorm:"type:varchar(100); not null"`
	Stmpport      int    `gorm:"type:"`
	Emailusername string `gorm:"type:"`
	Emailpassword string `gorm:"type:"`
}
