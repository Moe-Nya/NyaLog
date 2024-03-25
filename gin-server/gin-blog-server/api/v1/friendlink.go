package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增友链
func CreateFriendLink(c *gin.Context) {
	var data model.Friendlink
	_ = c.ShouldBindJSON(&data)
	err := service.CreateFriendLink(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "new friend link success",
		})
	}
}

// 查询友链列表
func SeleFriendLink(c *gin.Context) {
	var pagedata service.FriendLinkPage
	var friendlinks []model.Friendlink
	_ = c.ShouldBindJSON(&pagedata)
	friendlinks, err, total := service.SeleFriendLink(&pagedata)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        err,
			"total":       total,
			"friendlinks": friendlinks,
		})
	}
}

// 修改友链
func ModifyFriendLink(c *gin.Context) {
	var data model.Friendlink
	_ = c.ShouldBindJSON(&data)
	err := service.ModifyFriendLink(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "modify friend link success",
		})
	}
}

// 删除友链
func DeleFriendLink(c *gin.Context) {
	var data model.Friendlink
	_ = c.ShouldBindJSON(&data)
	err := service.DeleFriendLink(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "delete friend link success",
		})
	}
}
