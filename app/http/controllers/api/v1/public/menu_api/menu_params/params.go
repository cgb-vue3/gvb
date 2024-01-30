package menu_params

// imageIdList 传入的图片ad列表
type imageIdList struct {
	ID uint `json:"id"`
}

// AddMenuParams 添加菜单的参数
type AddMenuParams struct {
	Title        string        `json:"title" binding:"required"`     // 标题
	TitleEn      string        `json:"titleEn" binding:"required"`   // 标题英文
	Slogan       string        `json:"slogan"`                       // 简介
	ImageCutTime int           `json:"image_cut_time"`               // 轮播时间
	ImageList    []imageIdList `json:"image_list"`                   // 图片ID
	Sort         int           `json:"sort" binding:"required,gt=1"` // 菜单排序
}

type PutMenuParams struct {
	ID           uint          `json:"id" binding:"required"`        // 菜单id
	Title        string        `json:"title" binding:"required"`     // 标题
	TitleEn      string        `json:"titleEn" binding:"required"`   // 标题英文
	Slogan       string        `json:"slogan"`                       // 简介
	ImageCutTime int           `json:"image_cut_time"`               // 轮播时间
	ImageList    []imageIdList `json:"image_list"`                   // 图片ID
	Sort         int           `json:"sort" binding:"required,gt=1"` // 菜单排序
}

// DeleteMenuParams 需要删除的菜单id
type DeleteMenuParams struct {
	IDList []uint `json:"id_list" binding:"required"` // 菜单id
}
