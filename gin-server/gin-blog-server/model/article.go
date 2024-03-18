package model

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Articleid    int64  `gorm:"type:bigint;not null;primary_key" json:"articleid" label:"文章id"`
	Articleimg   string `gorm:"type:varchar(1000)" json:"articleimg" label:"文章头图"`
	Articletitle string `gorm:"type:text;not null" json:"articletitle" label:"文章标题"`
	Articlelikes string `gorm:"type:varchar(100);not null;default:0" json:"articlelikes" label:"文章点赞数"`
	Articleviews string `gorm:"type:varchar(100);not null;default:0" json:"articleviews" label:"文章浏览数"`
	Cid          int    `gorm:"type:int(5)" json:"cid" label:"文章分类id"`
	Aiswitch     int    `gorm:"type:int(5);not null;default:0" json:"aiswitch" label:"谋篇文章AI摘要开关"`
	Aisummary    string `gorm:"type:text" json:"aisummary" label:"AI文章摘要"`
	Text         string `gorm:"type:text;not null" json:"text" label:"文章内容"`
	Shorttext    string `gorm:"type:text;not null" json:"shorttext" label:"短原文"`
}

// 新增文章
func CreateArticle(data *Article) int {
	var article Article
	article.Articleimg = data.Articleimg
	article.Articletitle = data.Articletitle
	article.Articlelikes = data.Articlelikes
	article.Articleviews = data.Articleviews
	article.Cid = data.Cid
	article.Aiswitch = data.Aiswitch
	article.Aisummary = data.Aisummary
	article.Text = data.Text
	// 短原文截取原文的前30字符
	// 这样做在列出文章列表是，可以丰富展示框内容，又节约服务器性能
	// 只截取30字符，交给前端取舍
	runes := []rune(data.Text)
	if len(runes) > 30 {
		runes = runes[:30]
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
	if err != nil {
		return article, errmsg.ERROR
	}
	newViews := utils.BigNumAdd(article.Articleviews)
	err = db.Model(&article).Update("articleviews", newViews).Error
	if err != nil {
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
	return articleList, errmsg.SUCCESS, total
}

// 编辑文章
func ModifyArticle(data *Article) int {
	var article Article
	var articlemaps = make(map[string]interface{})
	articlemaps["articleimg"] = data.Articleimg
	articlemaps["articletitle"] = data.Articletitle
	articlemaps["cid"] = data.Cid
	articlemaps["aiswitch"] = data.Aiswitch
	articlemaps["aisummary"] = data.Aisummary
	articlemaps["text"] = data.Text
	runes := []rune(data.Text)
	if len(runes) > 30 {
		runes = runes[:30]
	}
	articlemaps["shorttext"] = string(runes)
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
func AddLike(articleid int) int {
	var article Article
	err := db.Where("articleid = ?", articleid).Find(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	articlelikes := article.Articlelikes
	newlikes := utils.BigNumAdd(articlelikes)
	err = db.Model(&article).Update("articlelikes", newlikes).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 文章CID更新(如果数据库无法关联，就手动)
func DeleCid(cid int) int {
	var article Article
	var newcid *int = nil
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
