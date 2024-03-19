package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------------------------------------------
// 创建页面
func CreatePage(c *gin.Context) {
	var data model.Page
	_ = c.ShouldBindJSON(&data)
	err := service.CreatePage(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "page create success",
		})
	}
}

// -------------------------------------------
// 查询页面
func SelePage(c *gin.Context) {
	purl := c.Param("purl")
	var page model.Page
	page, err := service.SelePage(purl)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": err,
			"page": page,
		})
	}
}

// 查询页面列表
func SelectPageList(c *gin.Context) {
	var pagedata service.PageList
	_ = c.ShouldBindJSON(&pagedata)
	var pages []model.Page
	pages, err, total := service.SelectPageList(&pagedata)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  err,
			"total": total,
			"pages": pages,
		})
	}
}

// -------------------------------------------
// 编辑页面
func EditPage(c *gin.Context) {
	var data model.Page
	_ = c.ShouldBindJSON(&data)
	err := service.EditPage(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "edit page success",
		})
	}
}

// -------------------------------------------
// 删除页面
func DeletePage(c *gin.Context) {
	var data model.Page
	_ = c.ShouldBindJSON(&data)
	err := service.DeletePage(data.Pid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "delete page success",
		})
	}
}
