package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"encoding/xml"
	"strconv"
	"time"
)

type Rss struct {
	XMLName  xml.Name `xml:"rss"`
	Version  string   `xml:"version,attr"`
	Articles Channel  `xml:"channel"`
}

type Channel struct {
	Title       string     `xml:"title"`
	Description string     `xml:"description"`
	Link        string     `xml:"link"`
	PubDate     string     `xml:"pubDate"`
	Content     []Articles `xml:"item"`
}

type Articles struct {
	Title   string    `xml:"title"`
	Link    string    `xml:"link"`
	Content string    `xml:"content"`
	PubDate time.Time `xml:"pubDate"`
}

// 获取前10篇文章内容
func GetArticle() []Articles {
	var articles []model.Article
	articles, _ = model.RssArtuckeList()
	var articleInfo []Articles
	for _, i := range articles {
		var a Articles
		a.Title = i.Articletitle
		a.Link = utils.Domain + "/article/" + strconv.FormatInt(i.Articleid, 10)
		a.Content = i.Text
		a.PubDate = i.CreatedAt
		articleInfo = append(articleInfo, a)
	}
	return articleInfo
}

func GenerateRSS() Rss {
	articleInfo := GetArticle()
	var rss Rss
	var channel Channel
	rss.Version = "2.0"
	channel.Title = utils.Author
	channel.Description = utils.Description
	channel.Link = utils.Domain
	channel.PubDate = time.Now().Format(time.RFC1123)
	channel.Content = articleInfo
	rss.Articles = channel
	return rss
}
