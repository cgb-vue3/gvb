package menu_resp

import model "akita/app/http/models"

// RespImageList 存放与菜单相关的图片信息
type RespImageList struct {
	ID   uint
	Name string
	Path string
}

// RespMenuList 用于返回菜单列表
type RespMenuList struct {
	MenuModel model.MenuModel
	Images    []RespImageList
}

// RespDeleteMenuInfoList 返回删除菜单的信息
type RespDeleteMenuInfoList struct {
	ID  uint   `json:"id"`
	Msg string `json:"msg"`
}
