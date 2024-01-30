package comment_manage_params

// CommentParams 添加评论参数
type CommentParams struct {
	ArticleID    uint   `json:"article_id"`     // 评论文章的ID
	ParCommentID uint   `json:"par_comment_id"` // 父评论ID
	ParUserID    uint   `json:"par_user_id"`    // 父评论用户id
	SelfUserID   uint   `json:"self_user_id"`   // 子评论用户的ID
	Content      string `json:"content"`        // 评论内容
	Deep         int    `json:"deep"`           // 评论深度
}

// Get 获取评论参数
type Get struct {
	ArticleID uint `form:"article_id"` // 评论文章的ID
	Deep      int  `json:"deep"`       // 评论深度
}
