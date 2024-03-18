package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
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
