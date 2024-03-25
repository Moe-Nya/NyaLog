package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 新增友链
func CreateFriendLink(data *model.Friendlink) int {
	err := model.CreateFriendLink(data)
	if err != errmsg.SUCCESS {
		return errmsg.CreateFriendLinkFailed
	}
	return errmsg.SUCCESS
}

// 友链分页结构体
type FriendLinkPage struct {
	PageSize int `json:"pagesize"`
	PageNum  int `json:"pagenum"`
}

// 查询友链列表
func SeleFriendLink(pageData *FriendLinkPage) ([]model.Friendlink, int, int64) {
	var friendlinks []model.Friendlink
	friendlinks, err, total := model.SeleFriendLink(pageData.PageSize, pageData.PageNum)
	if err != errmsg.SUCCESS {
		return friendlinks, errmsg.FriendLinkQueryFailed, 0
	}
	return friendlinks, errmsg.SUCCESS, total
}

// 修改友链
func ModifyFriendLink(data *model.Friendlink) int {
	err := model.ModifyFriendLink(data)
	if err != errmsg.SUCCESS {
		return errmsg.FriendLinkModifyFailed
	}
	return errmsg.SUCCESS
}

// 删除友链
func DeleFriendLink(data *model.Friendlink) int {
	err := model.DeleFriendLink(data.Friendlinkid)
	if err != errmsg.SUCCESS {
		return errmsg.FriendLinkDeleteFailed
	}
	return errmsg.SUCCESS
}
