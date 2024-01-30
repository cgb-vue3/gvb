package article_manage_resp

import "time"

// RespList 获取文章列表返回信息
type RespList struct {
	ID           uint      `json:"id"`           // 文章id
	Author       string    `json:"author"`       // 作者
	AuthorAvatar string    `json:"authorAvatar"` // 作者头像
	Title        string    `json:"title"`        // 文章标题
	Abstract     string    `json:"abstract"`     //文章简介
	Content      string    `json:"content"`      // 文章内容
	LookCount    int       `json:"lookCount"`    // 文章浏览量
	CommentCount int       `json:"commentCount"` // 文章评论量
	LikeCount    int       `json:"likeCount"`    // 文章点赞量
	Tags         []string  `json:"tags"`         // 文章关联的标签
	Banner       string    `json:"banner"`       // 文章封面
	Category     string    `json:"category"`     // 文章分类
	Issue        time.Time `json:"issue"`        // 发布时间
}
