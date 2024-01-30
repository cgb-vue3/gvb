package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

// Role 角色的类型
type Role int

const (
	PermissionSuperAdmin  Role = iota + 1 // 超级管理员:拥有所有权限
	PermissionAdin                        // 管理员：拥有部分权限
	PermissionUser                        // 普通用户：没有管理权限
	PermissionVisitor                     // 游客：没有权限，只能浏览
	PermissionDisableUser                 // 被禁用的用户：禁止登录
)

// MarshalJSON json序列化Permission
func (R Role) MarshalJSON() ([]byte, error) {
	var str string
	switch R {
	case PermissionSuperAdmin:
		str = "超级管理员"
	case PermissionAdin:
		str = "管理员"
	case PermissionUser:
		str = "普通用户"
	case PermissionVisitor:
		str = "游客"
	case PermissionDisableUser:
		str = "被禁言"
	default:
		str = "其它"
	}
	return json.Marshal(str)
}

// UserModel 用户表结构
type UserModel struct {
	gorm.Model
	Nickname string `json:"nickname"`               // 昵称
	UserName string `json:"username"`               // 用户名
	Slogan   string `json:"slogan"`                 // 个签
	PassWord string `json:"-"`                      // 密码
	Sex      string `json:"sex"`                    // 性别
	Addr     string `json:"addr"`                   // 地址
	Avatar   string `json:"avatar"`                 // 头像
	Email    string `json:"email"`                  // 邮箱
	Phone    string `json:"phone"`                  // 手机号
	Wechat   string `json:"wechat"`                 // 微信号
	Token    string `json:"token"`                  // token
	Role     Role   `json:"role" gorm:"default:3" ` // 用户级别
	//ImageModel []ImageModel `gorm:"many2many:user_images"`
	ArticleModel []ArticleModel `json:"article_model"` // 发布的文章
}
