package child_menu_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type ChildMenuDao struct {
	Orm *gorm.DB
}

func NewChildMenuDao() ChildMenuDao {
	return ChildMenuDao{
		Orm: global.MDB,
	}
}
