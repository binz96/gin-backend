package model

import (
	"fmt"

	"github.com/binz96/blog/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey; auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

func CheckCategory(name string) int {
	var c Category
	db.Select("id").First(&c, "name = ?", name)
	fmt.Println(c)
	if c.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

func CreateCategory(c *Category) int {
	err = db.Create(&c).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetCategories(pageSize, pageNum int) []Category {
	var cates []Category
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cates
}

// TODO: 查询分类下的所有文章
// selecr article.title as 文章标题, article.content as 文章内容 category.name as 博客分类 from article
// inner join category
// on article.cid = category.id
// where category.id = 1
// GORM: 一个分类对应多篇文章，文章belongs to分类
// 物理外键有数据一致性，性能不好，现在很少用
func GetArticlesInCategory(cid string, pageSize, pageNum int) ([]Article, int) {
	var articles []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", cid).Find(&articles).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST
	}
	return articles, errmsg.SUCCESS
}

func UpdateCategory(id string, c *Category) int {
	m := make(map[string]interface{})
	m["name"] = c.Name
	err = db.Model(&Category{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCategory(id string) int {
	err = db.Delete(&Category{}, id).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
