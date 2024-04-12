package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 新增文章分类
func CreateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	err := service.CreateCategory(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "categoary create success",
		})
	}
}

// -------------------------------------------
// 查询文章分类
func SeleCategory(c *gin.Context) {
	var data []model.Category
	data, err := service.SeleCategory()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"category": data,
		})
	}
}

// 根据cid查询分类名
func SeleCatName(c *gin.Context) {
	var cid service.Catid
	_ = c.ShouldBindJSON(&cid)
	ctext, err := service.SeleCatName(cid.Cid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  err,
			"ctext": ctext,
		})
	}
}

// -------------------------------------------
// // 编辑文章分类
func EditCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	err := service.EditCategory(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "category edit success",
		})
	}
}

// -------------------------------------------
// 删除文章分类
func DeleteCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	err := service.DeleteCategory(data.Cid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "category delete success",
		})
	}
}
