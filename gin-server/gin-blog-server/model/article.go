package model

import (
	"NyaLog/gin-blog-server/utils/errmsg"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Articleid    int64      `gorm:"type:bigint;not null;primary key" json:"articleid" label:"文章id"`
	Articleimg   string     `gorm:"type:varchar(1000)" json:"articleimg" label:"文章头图"`
	Articletitle string     `gorm:"type:text;not null" json:"articletitle" label:"文章标题"`
	Articledate  time.Time  `gorm:"type:datetime;not null" json:"articledate" label:"文章最后更新时间"`
	Articlelikes string     `gorm:"type:varchar(100);not null;default:0" json:"articlelikes" label:"文章点赞数"`
	Articleviews string     `gorm:"type:varchar(100)not null;default:0" json:"articleviews" label:"文章浏览数"`
	Cid          []Category `gorm:"type:int;foreignKey:Cid" json:"cid" label:"文章分类id"`
	Aisummary    string     `gorm:"type:text" json:"aisummary" label:"AI文章摘要"`
	Text         string     `gorm:"type:text;not null" json:"text" label:"文章内容"`
}

// 新增文章
func CreateArticle(data *Article) int {
	var article Article
	article.Articleimg = data.Articleimg
	article.Articletitle = data.Articletitle
	// 若输入的时间为空，则使用当前时间作为文章发布时间
	if data.Articledate.IsZero() {
		data.Articledate = time.Now()
	}
	article.Articledate = data.Articledate
	article.Articlelikes = data.Articlelikes
	article.Articleviews = data.Articleviews
	article.Cid = data.Cid
	article.Aisummary = data.Aisummary
	article.Text = data.Text
	var count int64
	var a Article
	err := db.Find(&a).Count(&count).Error
	if err != nil {
		return errmsg.ERROR
	}
	// 使navid有一个自增效果，方便后续的寻找和修改
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
func SeleOneArticle(articleid int) (Article, int) {
	var article Article
	err := db.Where("articleid = ?", articleid).Find(&article).Error
	if err != nil {
		return article, errmsg.ERROR
	}

	return article, errmsg.SUCCESS
}

// 查询文章列表

// 编辑文章

// 删除文章
