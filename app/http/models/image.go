package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

// ImgType 自定义一个类型
type ImgType int

// 定义常量
const (
	local ImgType = iota
	qiNiu
)

// MarshalJSON 通过switch判断，用一个变量去接，序列化这个变量
func (Img ImgType) MarshalJSON() ([]byte, error) {
	var str string
	switch Img {
	case local:
		str = "本地"
	case qiNiu:
		str = "七牛云"
	}
	return json.Marshal(str)
}

// ImageModel 图片表结构
type ImageModel struct {
	gorm.Model
	Name string  `json:"name"`                 // 文件名
	Path string  `json:"path"`                 // 图片路径
	Hash string  `json:"hash"`                 // hash加密
	Env  ImgType `json:"env" gorm:"default:0"` // 本地或七牛
	Type string  `json:"type"`                 // 图片类型
	//ArticleModel   ArticleModel `json:"articleModel" gorm:"foreignKey:ArticleModelID"` // 文章
	//ArticleModelID uint         `json:"articleModelID"`                                // 文章ID
	//UserModel      UserModel    `json:"user_model" gorm:"foreignKey:UserModelID"`      //文章作者
	//UserModelID    uint         `json:"user_id"`                                       // 文章作者ID
}
