package model

import (
	"github.com/binz96/blog/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100); not null" json:"title"`
	Cid     int    `gorm:"type:int; not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

func CreateArticle(a *Article) int {
	err = db.Create(&a).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func GetArticle(id string) (Article, int) {
	var a Article
	err = db.Preload("Category").Where("id = ?", id).First(&a).Error
	if err != nil {
		return a, errmsg.ERROR_ART_NOT_EXIST
	}
	return a, errmsg.SUCCESS
}

func GetArticles(pageSize, pageNum int) ([]Article, int) {
	var articles []Article
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}

func UpdateArticle(id string, a *Article) int {
	m := make(map[string]interface{})
	m["title"] = a.Title
	m["cid"] = a.Cid
	m["desc"] = a.Desc
	m["content"] = a.Content
	m["img"] = a.Img
	err = db.Model(&Article{}).Where("id = ?", id).Updates(m).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteArticle(id string) int {
	err = db.Delete(&Article{}, id).Error
	// DELETE FROM users WHERE id = 10;
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
