package models

import "gorm.io/gorm"

// TagModel 标签表结构
type TagModel struct {
	gorm.Model
	Title string `json:"title"` // 标签名
	//Articles []ArticleModel `json:"articles" gorm:"many2many:article_tag"` // 标签关联的文章
}

type a struct {
}
