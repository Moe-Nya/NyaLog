package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 新增Findme
func CreateFindme(data *model.FindMe) int {
	err := model.CreateFindme(data)
	if err != errmsg.SUCCESS {
		return errmsg.CreateFindMeFailed
	}
	return errmsg.SUCCESS
}

// 查询所有Findme
func SeleFindme() ([]model.FindMe, int) {
	var findme []model.FindMe
	findme, err := model.SeleFindme()
	if err != errmsg.SUCCESS {
		return findme, errmsg.FindMeQueryFailed
	}
	return findme, errmsg.SUCCESS
}

// 修改Findme
func ModifyFindme(data *model.FindMe) int {
	err := model.ModifyFindme(data)
	if err != errmsg.SUCCESS {
		return errmsg.FindMeModifyFailed
	}
	return errmsg.SUCCESS
}

// 删除Findme
func DeleFindme(data *model.FindMe) int {
	err := model.DeleFindme(data.FindId)
	if err != errmsg.SUCCESS {
		return errmsg.DeleteFindMeFailed
	}
	return errmsg.SUCCESS
}
