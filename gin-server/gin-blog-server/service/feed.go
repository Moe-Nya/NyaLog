package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils"
	"bytes"
	"encoding/xml"
	"strconv"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Content     string    `xml:"content"`
	PubDate     time.Time `xml:"pubDate"`
}

// markdown to html
func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
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
		var article = []byte(i.Text)
		var html = mdToHTML(article)
		var rune = bytes.Runes(html)
		a.Content = string(rune)
		var description = []byte(i.Shorttext)
		var htmlDes = mdToHTML(description)
		var runeDes = bytes.Runes(htmlDes)
		a.Description = string(runeDes)
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
