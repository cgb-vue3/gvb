package common

// PagingParams 返回图片列表的参数
type PagingParams struct {
	Page     int    `form:"page" json:"page" binding:"required,gte=1"`           // 页数
	PageSize int    `form:"page_size" json:"page_size" binding:"required,gte=1"` // 条数
	Sort     int    `form:"sort" json:"sort" binding:"oneof=0 1"`                // 排序 0：正序、1：反序
	Type     string `form:"type" json:"type"`
}
