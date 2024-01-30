package models

import "gorm.io/gorm"

// ArticleModel 文章表结构
type ArticleModel struct {
	gorm.Model
	Title        string     `json:"title"`                             // 文章标题
	Abstract     string     `json:"abstract"`                          //文章简介
	Content      string     `json:"content"`                           // 文章内容
	LookCount    int        `json:"lookCount"`                         // 文章浏览量
	CommentCount int        `json:"commentCount"`                      // 文章评论量
	LikeCount    int        `json:"likeCount"`                         // 文章点赞量
	TagModel     []TagModel `json:"tags" gorm:"many2many:article_tag"` // 文章关联的标签
	Banner       string     `json:"banner"`                            //文章封面
	//CommentModel []CommentModel `json:"comments"`                          // 文章关联的评论列表
	UserModel       UserModel     `json:"-" gorm:"foreignKey:UserModelID"` //文章关联的作者
	UserModelID     uint          `json:"userModelID"`                     // 文章作者ID
	CategoryModel   CategoryModel `json:"categoryModel"`                   // 文章关联分类
	CategoryModelID uint          `json:"categoryModelID"`                 // 文章分类ID
	//ImageModel []ImageModel `json:"image_model"` // 文章关联的封面
}
