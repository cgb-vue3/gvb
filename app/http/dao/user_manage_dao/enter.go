package user_manage_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type UserManageDao struct {
	Orm *gorm.DB
}

func NewUserManageDao() UserManageDao {
	return UserManageDao{
		Orm: global.MDB,
	}
}
