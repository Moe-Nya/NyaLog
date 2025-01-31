package model

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Articleid       int64  `gorm:"type:bigint;not null;primary_key" json:"articleid" label:"文章id"`
	Articleimg      string `gorm:"type:varchar(1000)" json:"articleimg" label:"文章头图"`
	Articletitle    string `gorm:"type:text;not null" json:"articletitle" label:"文章标题"`
	Articlelikes    string `gorm:"type:varchar(100);default:0" json:"articlelikes" label:"文章点赞数"`
	Articleviews    string `gorm:"type:varchar(100);default:0" json:"articleviews" label:"文章浏览数"`
	Cid             int    `gorm:"type:int(5);primary_key" json:"cid" label:"文章分类id"`
	Aiswitch        int    `gorm:"type:int(5);not null;default:0" json:"aiswitch" label:"谋篇文章AI摘要开关"`
	Aisummary       string `gorm:"type:text" json:"aisummary" label:"AI文章摘要"`
	Text            string `gorm:"type:text;not null" json:"text" label:"文章内容"`
	Shorttext       string `gorm:"type:text;not null" json:"shorttext" label:"短原文"`
	Articlecategory int    `gorm:"type:int(5);not null" json:"articlecategory" label:"文章类型(站内or站外)"`
}

// 新增文章
func CreateArticle(data *Article) int {
	if data.Articletitle == "" || data.Text == "" {
		return errmsg.ERROR
	}
	var article Article
	article.Articleimg = data.Articleimg
	article.Articletitle = data.Articletitle
	article.Articlelikes = "0"
	article.Articleviews = "0"
	article.Cid = data.Cid
	article.Aiswitch = data.Aiswitch
	article.Aisummary = data.Aisummary
	article.Text = data.Text
	article.Articlecategory = data.Articlecategory
	// 如果是站外文章，不开启AI摘要 0是站内 1是站外
	if article.Articlecategory == 1 {
		article.Aiswitch = 0
	}
	// 短原文截取原文的前30字符
	// 这样做在列出文章列表是，可以丰富展示框内容，又节约服务器性能
	// 只截取30字符，交给前端取舍
	runes := []rune(data.Text)
	if len(runes) > 100 {
		runes = runes[:100]
	}
	article.Shorttext = string(runes)

	// 使id有一个自增效果，方便后续的寻找和修改
	var count int64
	var a Article
	err := db.Find(&a).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	if count == 0 {
		article.Articleid = 0
	} else {
		var finalarticle Article
		err = db.Limit(1).Offset(int(count - 1)).Find(&finalarticle).Error
		if err != nil {
			return errmsg.ERROR
		}
		article.Articleid = finalarticle.Articleid + 1
	}
	err = db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询单篇文章
func SeleOneArticle(articleid int64) (Article, int) {
	var article Article
	err := db.Where("articleid = ?", articleid).Find(&article).Error
	if err != nil || article.Articletitle == "" {
		return article, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}

// 查询文章列表
func SeleListArticle(pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	err := db.Select("article.articleid, articleimg, articletitle, created_at, updated_at, articlelikes, articleviews, shorttext, cid").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Find(&articleList).Error
	if err != nil {
		return articleList, errmsg.ERROR, 0
	}
	var total int64
	db.Model(&articleList).Count(&total)
	for i := range articleList {
		articleList[i].Text = ""
	}
	return articleList, errmsg.SUCCESS, total
}

// 增加文章浏览量
func AddViews(articleid int64, views string) int {
	article, err := SeleOneArticle(articleid)
	if err == errmsg.ERROR {
		return errmsg.ERROR
	}
	newViews := utils.BigDecimal(article.Articleviews, views)
	e := db.Model(&article).Update("articleviews", newViews).Error
	if e != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// RSS查询前10篇文章
func RssArtuckeList() ([]Article, int) {
	var articleList []Article
	err := db.Select("article.articleid, articletitle, created_at, shorttext, text").Limit(10).Order("Created_At DESC").Find(&articleList).Error
	if err != nil {
		return articleList, errmsg.ERROR
	}
	return articleList, errmsg.SUCCESS
}

type Archive struct {
	Year     int       `json:"year"`
	Articles []Article `json:"articles"`
}

// 查询文章归档
func SeleArchive() ([]Archive, int) {
	var archives []Archive
	var years []int
	err := db.Table("article").Select("YEAR(created_at)").Group("YEAR(created_at)").Order("YEAR(created_at) DESC").Scan(&years).Error
	if err != nil {
		return archives, errmsg.ERROR
	}
	for i := range years {
		var articles []Article
		var archive Archive
		_ = db.Table("article").Select("article.articleid, articletitle, created_at").Where("YEAR(created_at)=?", years[i]).Order("created_at DESC").Scan(&articles).Error
		archive.Year = years[i]
		archive.Articles = articles
		archives = append(archives, archive)
	}
	return archives, errmsg.SUCCESS
}

// 编辑文章
func ModifyArticle(data *Article) int {
	if data.Articletitle == "" || data.Text == "" {
		return errmsg.ERROR
	}
	var article Article
	var articlemaps = make(map[string]interface{})
	articlemaps["articleimg"] = data.Articleimg
	articlemaps["articletitle"] = data.Articletitle
	articlemaps["cid"] = data.Cid
	articlemaps["aiswitch"] = data.Aiswitch
	articlemaps["aisummary"] = data.Aisummary
	articlemaps["text"] = data.Text
	articlemaps["articlecategory"] = data.Articlecategory

	// 如果是站外文章，不开启AI摘要 0是站内 1是站外
	if data.Articlecategory == 1 {
		articlemaps["aiswitch"] = 0
		articlemaps["aisummary"] = ""
	}

	runes := []rune(data.Text)
	if len(runes) > 50 {
		runes = runes[:50]
	}
	articlemaps["shorttext"] = string(runes)
	if data.Articletitle == "" {
		return errmsg.ERROR
	}
	err := db.Model(&article).Where("articleid = ?", data.Articleid).Updates(articlemaps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(articleid int64) int {
	var article Article
	err := db.Where("articleid = ?", articleid).Unscoped().Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 文章喜欢数累加
func AddLike(articleid int64, likes string) int {
	var article Article
	err := db.Where("articleid = ?", articleid).Find(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	if article.Articletitle == "" {
		return errmsg.ERROR
	}
	newlikes := utils.BigDecimal(article.Articlelikes, likes)
	err = db.Model(&article).Update("articlelikes", newlikes).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 文章CID更新(如果数据库无法关联，就手动)
func DeleCid(cid int) int {
	var article Article
	var newcid int = -1
	err := db.Model(&article).Where("cid = ?", cid).Update("cid", newcid).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 同CID下的文章列表
func SameCidArt(cid int) ([]Article, int) {
	var article []Article
	err := db.Where("cid = ?", cid).Find(&article).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return article, errmsg.SUCCESS
}
