package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 新增页面
func CreatePage(data *model.Page) int {
	err := model.CreatePage(data)
	if err != errmsg.SUCCESS {
		return errmsg.PageCreateFailed
	}
	return errmsg.SUCCESS
}

// 查询页面
func SelePage(purl string) (model.Page, int) {
	var page model.Page
	page, err := model.SelePage(purl)
	if err != errmsg.SUCCESS {
		return page, errmsg.PageQueryFailed
	}
	return page, errmsg.SUCCESS
}

// 查询文章列表就所需要的结构体
type PageList struct {
	PageSize int `json:"pagesize"`
	PageNum  int `json:"pagenum"`
}

// 查询页面列表
func SelectPageList(pagedata *PageList) ([]model.Page, int, int64) {
	var pagelist []model.Page
	pagelist, err, total := model.SelectPageList(pagedata.PageSize, pagedata.PageNum)
	if err != errmsg.SUCCESS {
		return pagelist, errmsg.PageQueryFailed, total
	}
	return pagelist, errmsg.SUCCESS, total
}

// 编辑页面
func EditPage(data *model.Page) int {
	err := model.EditPage(data)
	if err != errmsg.SUCCESS {
		return errmsg.PageEditFailed
	}
	return errmsg.SUCCESS
}

// 删除页面
func DeletePage(pid int) int {
	err := model.DeletePage(pid)
	if err != errmsg.SUCCESS {
		return errmsg.PageDeleteFailed
	}
	return errmsg.SUCCESS
}
