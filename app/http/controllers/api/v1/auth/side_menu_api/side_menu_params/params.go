package side_menu_params

// SideMenuParams 侧栏菜单参数
type SideMenuParams struct {
	//ParentID  uint             `json:"parent_id,omitempty"` // 父级id
	Title string `json:"title"` // 菜单标题
	Icon  string `json:"icon"`  // 图标
	Name  string `json:"name"`  // 路由名
	Path  string `json:"path"`  // 路由路径
	Sort  int    `json:"sort"`  // 排序
	Level int    `json:"level"` // 用于判断菜单权限
}
