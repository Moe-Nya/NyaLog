package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 生成AI摘要
func CreateAISummary(text string, blogset *model.BlogSet) string {
	if blogset.Aicategory == 0 {
		AisummaryData := middleware.GPT(utils.AIsummaryURL, utils.AIsummaryKEY, "gpt-3.5-turbo", text, "Chinese")
		aisummary := AisummaryData.Text
		return aisummary
	} else if blogset.Aicategory == 1 {
		AisummaryData := middleware.QianWen(utils.AIsummaryURL, utils.AIsummaryKEY, "qwen-turbo", text, "Chinese")
		aisummary := AisummaryData.Text
		return aisummary
	}
	return ""
}

// 新增文章
func CreateArticle(data *model.Article) int {
	if data.Aiswitch == 1 {
		var blogset model.BlogSet
		blogset, err := model.SeleBlogSet()
		if err != errmsg.SUCCESS {
			return errmsg.BlogNotSet
		}
		if blogset.Aiswitch == 1 {
			aisummary := CreateAISummary(data.Text, &blogset)
			if aisummary != "" {
				data.Aisummary = aisummary
			}
		}
	}
	err := model.CreateArticle(data)
	if err != errmsg.SUCCESS {
		return errmsg.ArticleUpdateFailed
	}
	return errmsg.SUCCESS
}

// 编辑文章
func EditArticle(data *model.Article) int {
	if data.Aiswitch == 1 {
		var blogset model.BlogSet
		blogset, err := model.SeleBlogSet()
		if err != errmsg.SUCCESS {
			return errmsg.BlogNotSet
		}
		if blogset.Aiswitch == 1 {
			aisummary := CreateAISummary(data.Text, &blogset)
			if aisummary != "" {
				data.Aisummary = aisummary
			}
		}
	}
	err := model.ModifyArticle(data)
	if err != errmsg.SUCCESS {
		return errmsg.ArticleUpdateFailed
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(articleid int) int {
	err := model.DeleteArticle(articleid)
	if err != errmsg.SUCCESS {
		return errmsg.ArticleUpdateFailed
	}
	return errmsg.SUCCESS
}

// 查询单篇文章
func SeleOneArticle(articleid int) (model.Article, int) {
	var article model.Article
	article, err := model.SeleOneArticle(articleid)
	if err != errmsg.SUCCESS {
		return article, errmsg.ArticleQueryFailed
	}
	return article, errmsg.SUCCESS
}

// 查询文章列表
