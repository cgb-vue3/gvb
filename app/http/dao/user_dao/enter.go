package user_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type UserDao struct {
	Orm *gorm.DB
}

func NewUserDao() UserDao {
	return UserDao{
		Orm: global.MDB,
	}
}
