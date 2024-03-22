package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 增加导航标签
func CreateNav(data *model.Navigation) int {
	err := model.CreateNav(data)
	if err != errmsg.SUCCESS {
		return errmsg.CreateNavFailed
	}
	return errmsg.SUCCESS
}

// 查询所有导航标签
func SeleNav() ([]model.Navigation, int) {
	var navs []model.Navigation
	navs, err := model.SeleNav()
	if err != errmsg.SUCCESS {
		return navs, errmsg.NavQueryFailed
	}
	return navs, errmsg.SUCCESS
}

// 修改导航标签
func ModifyNav(data *model.Navigation) int {
	err := model.ModifyNav(data)
	if err != errmsg.SUCCESS {
		return errmsg.NavEditFailed
	}
	return errmsg.SUCCESS
}

// 删除导航标签
func DeleNav(data *model.Navigation) int {
	err := model.DeleNav(data.NavId)
	if err != errmsg.SUCCESS {
		return errmsg.NavDeleteFailed
	}
	return errmsg.SUCCESS
}
