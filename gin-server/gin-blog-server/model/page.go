package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Pid      int    `gorm:"type:int(10);not null;primary_key" json:"pid" label:"页面编号"`
	Purl     string `gorm:"type:varchar(20);not null" json:"purl" label:"页面url地址"`
	Ptitle   string `gorm:"type:varchar(50);not null" json:"ptitle" label:"页面标题"`
	Pcontent string `gorm:"type:text" json:"pcontent" label:"内容"`
}

// 新增页面
func CreatePage(data *Page) int {
	var page Page
	page.Purl = data.Purl
	page.Ptitle = data.Ptitle
	page.Pcontent = data.Pcontent
	var count int64
	var p Page
	err := db.Find(&p).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// 使id有一个自增效果，方便后续的寻找和修改
	if count == 0 {
		p.Pid = 0
	} else {
		var finalpage Page
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalpage).Error
		if err != nil {
			return errmsg.ERROR
		}
		page.Pid = finalpage.Pid + 1
	}
	err = db.Create(&page).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询页面
func SelePage(purl string) (Page, int) {
	var page Page
	err := db.Where("purl = ?", purl).Find(&page).Error
	if err != nil {
		return page, errmsg.ERROR
	}
	return page, errmsg.SUCCESS
}

// 查询页面列表
func SelectPageList(pageSize int, pageNum int) ([]Page, int, int64) {
	var pages []Page
	err := db.Select("page.pid, purl, ptitle, created_at, updated_at, pcontent").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&pages).Error
	if err != nil {
		return pages, errmsg.ERROR, 0
	}
	var total int64
	db.Model(&pages).Count(&total)
	for i := range pages {
		pages[i].Pcontent = ""
	}
	return pages, errmsg.SUCCESS, total
}

// 编辑页面
func EditPage(data *Page) int {
	var page Page
	pagemaps := make(map[string]interface{})
	pagemaps["purl"] = data.Purl
	pagemaps["ptitle"] = data.Ptitle
	pagemaps["pcontent"] = data.Pcontent
	err := db.Model(&page).Where("pid = ?", data.Pid).Updates(&pagemaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除页面
func DeletePage(pid int) int {
	var page Page
	err := db.Where("pid = ?", pid).Unscoped().Delete(&page).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
