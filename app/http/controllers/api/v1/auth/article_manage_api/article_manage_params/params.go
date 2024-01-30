package article_manage_params

import "akita/app/http/controllers/common"

// AddParams 添加文章参数
type AddParams struct {
	ID            uint     `json:"id" binding:"required"`            // 用户ID
	Title         string   `json:"title" binding:"required"`         // 文章标题
	Content       string   `json:"content" binding:"required"`       // 文章内容
	Abstract      string   `json:"abstract" `                        //文章简介
	Banner        string   `json:"banner" binding:"required"`        //文章封面
	TagTitles     []string `json:"tagTitles"`                        // 文章标签
	CategoryTitle string   `json:"categoryTitle" binding:"required"` // 文章分类
	LookCount     int      `json:"lookCount"`                        // 文章浏览量
	CommentCount  int      `json:"commentCount"`                     // 文章评论量
	LikeCount     int      `json:"likeCount"`                        // 文章点赞量
}

// GetListParams 获取文章列表参数
type GetListParams struct {
	ID  uint                `json:"id" binding:"required"` // 用户ID
	Pag common.PagingParams `json:"pag"`                   // 分页参数
}

// PutParams 更新文章参数
type PutParams struct {
	UserID       uint   `json:"userID" binding:"required"`  // 用户ID
	Title        string `json:"title" binding:"required"`   // 文章标题
	Content      string `json:"content" binding:"required"` // 文章内容
	Abstract     string `json:"abstract" `                  //文章简介
	LookCount    int    `json:"lookCount"`                  // 文章浏览量
	CommentCount int    `json:"commentCount"`               // 文章评论量
	LikeCount    int    `json:"likeCount"`                  // 文章点赞量
}

// DelParams 删除文章参数
type DelParams struct {
	IDList []uint `json:"idList"`
}
