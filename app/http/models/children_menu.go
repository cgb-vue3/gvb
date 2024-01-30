package models

import "gorm.io/gorm"

// ChildSideMenuModel 侧栏菜单
type ChildSideMenuModel struct {
	gorm.Model
	Title           string `json:"title"` // 菜单标题
	Icon            string `json:"icon"`  // 图标
	Name            string `json:"name"`  // 路由名
	Path            string `json:"path"`  // 路由路径
	Sort            int    `json:"sort"`  // 排序
	Level           int    `json:"level"` // 菜单级别
	SideMenuModelID uint   `json:"sideMenuModelID"`
}
