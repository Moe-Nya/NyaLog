package service

import (
	"NyaLog/gin-blog-server/model"
	"NyaLog/gin-blog-server/utils/errmsg"
)

// 新增文章分类
func CreateCategory(data *model.Category) int {
	err := model.CreateCat(data)
	if err != errmsg.SUCCESS {
		return errmsg.CatCreateFailed
	}
	return errmsg.SUCCESS
}

// 查询文章分类
func SeleCategory() ([]model.Category, int) {
	var category []model.Category
	category, err := model.SeleCat()
	if err != errmsg.SUCCESS {
		return category, errmsg.CatQueryFailed
	}
	return category, errmsg.SUCCESS
}

// 编辑文章分类
func EditCategory(data *model.Category) int {
	err := model.ModifyCat(data)
	if err != errmsg.SUCCESS {
		return errmsg.CatUpdateFailed
	}
	return errmsg.SUCCESS
}

// 删除文章分类
func DeleteCategory(cid int) int {
	err := DeleCid(cid)
	if err != errmsg.SUCCESS {
		return errmsg.CatDeleteFailed
	}
	err = model.DeleCat(cid)
	if err != errmsg.SUCCESS {
		return errmsg.CatDeleteFailed
	}
	return errmsg.SUCCESS
}
