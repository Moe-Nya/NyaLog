package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增Findme
func CreateFindme(c *gin.Context) {
	var data model.FindMe
	_ = c.ShouldBindJSON(&data)
	err := service.CreateFindme(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "create findme success",
		})
	}
}

// 查询所有Findme
func SeleFindme(c *gin.Context) {
	var findme []model.FindMe
	findme, err := service.SeleFindme()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":   err,
			"findme": findme,
		})
	}
}

// 修改Findme
func ModifyFindme(c *gin.Context) {
	var data model.FindMe
	_ = c.ShouldBindJSON(&data)
	err := service.ModifyFindme(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "modify findme success",
		})
	}
}

// 删除Findme
func DeleFindme(c *gin.Context) {
	var data model.FindMe
	_ = c.ShouldBindJSON(&data)
	err := service.DeleFindme(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "delete findme success",
		})
	}
}
