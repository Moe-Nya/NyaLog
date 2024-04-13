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

// -------------------------------------------
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
	article, err, comswitch := service.SeleOneArticle(int64(id))
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":      err,
			"article":   article,
			"comswitch": comswitch,
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

// 查询文章归档
func SeleArchive(c *gin.Context) {
	var archives []model.Archive
	archives, err := service.SeleArchive()
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":     err,
			"archives": archives,
			"message":  "query archive success",
		})
	}
}

// -------------------------------------------
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

// -------------------------------------------
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
		// 删除文章的时候把对应的评论删掉
		_ = service.DeleteAllCom(article.Articleid)
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

// 文章喜欢数累加
func AddLike(c *gin.Context) {
	var articleid service.ArticleLike
	_ = c.ShouldBindJSON(&articleid)
	err := service.AddLike(articleid.Articleid)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": errmsg.GetErrorMsg(err),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    err,
			"message": "thx your like",
		})
	}
}
