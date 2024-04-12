package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Cid          int    `gorm:"type:int(5);not null;primary_key" json:"cid" label:"文章分类id"`
	Categorytext string `gorm:"type:varchar(50);not null" json:"categorytext" label:"分类内容"`
}

// 新增文章分类
func CreateCat(data *Category) int {
	if data.Categorytext == "" {
		return errmsg.ERROR
	}
	var count int64
	var c Category
	err := db.Find(&c).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// 使id有一个自增效果，方便后续的寻找和修改
	if count == 0 {
		data.Cid = 0
	} else {
		var finalcat Category
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalcat).Error
		if err != nil {
			return errmsg.ERROR
		}
		data.Cid = finalcat.Cid + 1
	}
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询文章分类列表
func SeleCat() ([]Category, int) {
	var cats []Category
	err := db.Find(&cats).Error
	if err != nil {
		return cats, errmsg.ERROR
	}
	return cats, errmsg.SUCCESS
}

// 根据cid查询分类名
func SeleCatName(cid int) (string, int) {
	var cat Category
	err := db.Select("category.categorytext").Where("cid=?", cid).First(&cat).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return "", errmsg.ERROR
	}
	return cat.Categorytext, errmsg.SUCCESS
}

// 编辑分类
func ModifyCat(data *Category) int {
	if data.Categorytext == "" {
		return errmsg.ERROR
	}
	var cat Category
	var catmap = make(map[string]interface{})
	catmap["categorytext"] = data.Categorytext
	var c Category
	err := db.First(&c, "cid = ?", data.Cid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Model(&cat).Where("cid = ?", data.Cid).Updates(catmap).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章分类
func DeleCat(cid int) int {
	var cat Category
	var c Category
	err := db.First(&c, "cid = ?", cid).Error
	if err != nil || err == gorm.ErrRecordNotFound {
		return errmsg.ERROR
	}
	err = db.Where("cid = ?", cid).Unscoped().Delete(&cat).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
