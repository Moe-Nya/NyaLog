package service

import (
	"NyaLog/gin-blog-server/middleware"
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"regexp"

	"github.com/cention-sany/html2text"
	"github.com/russross/blackfriday/v2"
)

// 生成AI摘要
func CreateAISummary(text string, blogset *model.BlogSet) string {
	html := blackfriday.Run([]byte(text))
	article, err := html2text.FromString(string(html))
	re := regexp.MustCompile(`[\n\r]+|\*\*|\*|__|_`)
	plainText := re.ReplaceAllString(article, "")
	if blogset.Aicategory == 0 {
		if err != nil {
			return ""
		}
		AisummaryData := middleware.GPT(utils.GPTURL, utils.GPTKey, "gpt-3.5-turbo", plainText, "Chinese")
		aisummary := AisummaryData.Text
		return aisummary
	} else if blogset.Aicategory == 1 {
		AisummaryData := middleware.QianWen(utils.QWURL, utils.QWKey, "qwen-turbo", plainText, "Chinese")
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
func DeleteArticle(articleid int64) int {
	err := model.DeleteArticle(articleid)
	if err != errmsg.SUCCESS {
		return errmsg.ArticleDeleteFailed
	}
	return errmsg.SUCCESS
}

// 删除某cid
func DeleCid(cid int) int {
	err := model.DeleCid(cid)
	if err != errmsg.SUCCESS {
		return errmsg.ArticleDeleteFailed
	}
	return errmsg.SUCCESS
}

// 查询单篇文章
func SeleOneArticle(articleid int64) (model.Article, int, int) {
	var article model.Article
	article, err := model.SeleOneArticle(articleid)
	if err != errmsg.SUCCESS {
		return article, errmsg.ArticleQueryFailed, 0
	}
	blogset, err := model.SeleBlogSet()
	if err != errmsg.SUCCESS {
		return article, errmsg.ArticleQueryFailed, 0
	}
	return article, errmsg.SUCCESS, blogset.Commentswitch
}

// 查询文章列表就所需要的结构体
type ArticleList struct {
	PageSize int `json:"pagesize"`
	PageNum  int `json:"pagenum"`
}

// 查询文章列表
func SeleListArticle(data *ArticleList) ([]model.Article, int, int64) {
	var articleList []model.Article
	articleList, err, total := model.SeleListArticle(data.PageSize, data.PageNum)
	if err != errmsg.SUCCESS {
		return articleList, errmsg.ArticleQueryFailed, total
	}
	return articleList, errmsg.SUCCESS, total
}

// 查询同cid文章列表
func SameCidArt(cid int) ([]model.Article, int) {
	var articleList []model.Article
	articleList, err := model.SameCidArt(cid)
	if err != errmsg.SUCCESS {
		return articleList, errmsg.ArticleQueryFailed
	}
	return articleList, errmsg.SUCCESS
}

type ArticleLike struct {
	Articleid int `json:"articleid"`
}

// 文章喜欢数累加
func AddLike(articleid int) int {
	err := model.AddLike(articleid)
	if err != errmsg.SUCCESS {

	}
	return errmsg.SUCCESS
}
