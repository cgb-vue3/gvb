package image_params

// TypeParams 返回图片列表的参数
type TypeParams struct {
	Type string `form:"type" json:"type"` // 上传的图片属于哪个功能
}

// PagingParams 返回图片列表的参数
type PagingParams struct {
	Page     int  `form:"page" binding:"required,gte=1"`      // 页数
	PageSize int  `form:"page_size" binding:"required,gte=1"` // 条数
	Sort     bool `form:"sort"`                               // 排序 true：正序 false：反序
	Type     TypeParams
}

// DeleteImgParams 删除图片的参数
type DeleteImgParams struct {
	IDList []uint `json:"id_list" binding:"required,min=1"`
}
