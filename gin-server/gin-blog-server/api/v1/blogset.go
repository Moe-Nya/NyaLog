package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 更新博客设置
func UpdateBlogSet(c *gin.Context) {
	var data model.BlogSet
	_ = c.ShouldBindJSON(&data)
	err := service.UpdateBlogSet(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "update success",
		})
	}
}

// 读取博客设置
func SeleBlogSet(c *gin.Context) {
	blogset, err := service.SeleBlogSet()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":           err,
			"sitename":       blogset.Sitename,
			"sitecreatetime": blogset.Sitecreatetime,
			"sitebackground": blogset.Sitebackground,
			"aiswitch":       blogset.Aiswitch,
			"aicategory":     blogset.Aicategory,
			"commentswitch":  blogset.Commentswitch,
		})
	}
}
