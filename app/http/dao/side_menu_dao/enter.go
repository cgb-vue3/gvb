package side_menu_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type SideMenuDao struct {
	Orm *gorm.DB
}

func NewSideMenuDao() SideMenuDao {
	return SideMenuDao{
		Orm: global.MDB,
	}
}
