package models

type ArticleTagModel struct {
	ArticleModel   ArticleModel `gorm:"foreignKey:ArticleModelID"`
	ArticleModelID uint         `json:"article_model_id"`
	TagModel       TagModel     `gorm:"foreignKey:TagModelID"`
	TagModelID     uint         `json:"tag_model_id"`
}
