package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 增加导航标签
func CreateNav(c *gin.Context) {
	var data model.Navigation
	_ = c.ShouldBindJSON(&data)
	err := service.CreateNav(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "create navigation success",
		})
	}
}

// 查询所有导航标签
func SeleNav(c *gin.Context) {
	var navs []model.Navigation
	navs, err := service.SeleNav()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":        err,
			"navigations": navs,
		})
	}
}

// 修改导航标签
func ModifyNav(c *gin.Context) {
	var data model.Navigation
	_ = c.ShouldBindJSON(&data)
	err := service.ModifyNav(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "edit navigation success",
		})
	}
}

// 删除导航标签
func DeleNav(c *gin.Context) {
	var data model.Navigation
	_ = c.ShouldBindJSON(&data)
	err := service.DeleNav(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "delete navigation success",
		})
	}
}
