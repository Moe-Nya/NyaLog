package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
	"time"
)

// 更新博客设置
func UpdateBlogSet(data *model.BlogSet) int {
	err, ex := model.ValidateBlogSet()
	if err != errmsg.SUCCESS {
		return errmsg.BlogSetFailed
	}
	if ex == 0 {
		e := model.CreateBlogSet(data)
		if e != errmsg.SUCCESS {
			return errmsg.BlogSetFailed
		}
	} else {
		e := model.ModifyBlogSet(data)
		if e != errmsg.SUCCESS {
			return errmsg.BlogSetFailed
		}
	}
	return errmsg.SUCCESS
}

// 读取博客设置
func SeleBlogSet() (model.BlogSet, int) {
	blogset, err := model.SeleBlogSet()
	if err != errmsg.SUCCESS {
		return blogset, errmsg.BlogNotSet
	}
	return blogset, errmsg.SUCCESS
}

// 博客设置公开信息结构体
type BlogSetInfo struct {
	Sitename       string    `json:"sitename"`
	Sitecreatetime time.Time `json:"sitecreatetime"`
	Sitebackground string    `json:"sitebackground"`
}

// 读取博客设置信息(公开)
func QueryBlogSet() (BlogSetInfo, int) {
	var blogsetinfo BlogSetInfo
	var blogset model.BlogSet
	blogset, err := model.QueryBlogSet()
	if err != errmsg.SUCCESS {
		return blogsetinfo, errmsg.QueryBlogSetFailed
	}
	blogsetinfo.Sitename = blogset.Sitename
	blogsetinfo.Sitecreatetime = blogset.Sitecreatetime
	blogsetinfo.Sitebackground = blogset.Sitebackground
	return blogsetinfo, errmsg.SUCCESS
}
