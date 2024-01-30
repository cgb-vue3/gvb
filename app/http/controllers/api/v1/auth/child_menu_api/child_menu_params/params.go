package child_menu_params

type ChildMenuParams struct {
	ParentID uint   `json:"parent_id,omitempty"` // 父级id
	Title    string `json:"title"`               // 菜单标题
	Icon     string `json:"icon"`                // 图标
	Name     string `json:"name"`                // 路由名
	Path     string `json:"path"`                // 路由路径
	Level    int    `json:"level"`               // 菜单级别
	Sort     int    `json:"sort"`                // 排序
}
