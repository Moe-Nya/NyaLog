package v1

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/service"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 新增文章
func CreateArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	err := service.CreateArticle(&article)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "article create success",
		})
	}
}

// 查询单篇文章
func SeleOneArticle(c *gin.Context) {
	idStr := c.Param("id")
	var article model.Article
	id, e := strconv.Atoi(idStr)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    errmsg.ArticleQueryFailed,
			"message": errmsg.GetErrorMsg(errmsg.ArticleQueryFailed),
		})
	}
	article, err := service.SeleOneArticle(int64(id))
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"article": article,
		})
	}
}

// 查询文章列表
func SeleListArticle(c *gin.Context) {
	var articleList []model.Article
	var data service.ArticleList
	_ = c.ShouldBindJSON(&data)
	articleList, err, total := service.SeleListArticle(&data)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"articles": articleList,
			"total":    total,
		})
	}
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	err := service.EditArticle(&article)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "article update success",
		})
	}
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	err := service.DeleteArticle(article.Articleid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "article delete success",
		})
	}
}

// 删除某CID
func DeleCid(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	err := service.DeleCid(article.Cid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "article delete success",
		})
	}
}
