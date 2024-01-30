package models

import "gorm.io/gorm"

// CommentModel 评论表结构
type CommentModel struct {
	gorm.Model
	ArticleModel   ArticleModel `json:"article" gorm:"foreignKey:ArticleModelID"` // 评论的文章
	ArticleModelID uint         `json:"articleModelID"`                           // 评论文章的ID
	ParCommentID   uint         `json:"par_comment_id"`                           // 父评论id
	ParUserID      uint         `json:"par_user_id"`                              // 父评论用户的id
	//SubUserID      uint         `json:"sub_user_id"`                              // 子评论的用户ID
	SelfID  uint   `json:"self_id"` // 评论用户的ID
	Content string `json:"content"` // 评论内容
	Deep    int    `json:"deep"`    // 评论深度
	//ParComments     *CommentModel   `json:"par_comments" gorm:"foreignKey:ParCommentID"` // 父评论
	//SubComments     []*CommentModel `json:"sub_comments" gorm:"foreignKey:ParCommentID"` // 子评论列表
}
