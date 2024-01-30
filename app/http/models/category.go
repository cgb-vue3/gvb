package models

import "gorm.io/gorm"

type CategoryModel struct {
	gorm.Model
	Title        string         `json:"title"`                      // 分类标题
	ArticleModel []ArticleModel `gorm:"foreignKey:CategoryModelID"` // 分类下的文章
	//ArticlesModel []ArticleModel
}
