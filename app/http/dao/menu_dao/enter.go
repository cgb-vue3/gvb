package menu_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type MenuDao struct {
	Orm *gorm.DB
}

func NewMenuDao() MenuDao {
	return MenuDao{
		Orm: global.MDB,
	}
}
