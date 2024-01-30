package models

import "gorm.io/gorm"

// MenuModel 菜单管理
type MenuModel struct {
	gorm.Model
	Title        string       `json:"title"`          // 标题
	TitleEn      string       `json:"titleEn"`        // 标题英文
	Slogan       string       `json:"slogan"`         // 简介
	ImageCutTime int          `json:"image_cut_time"` // 图片切换间隔
	ImageModel   []ImageModel `json:"-" gorm:"many2many:menu_images"`
	Sort         int          `json:"sort"` // 排序
}
