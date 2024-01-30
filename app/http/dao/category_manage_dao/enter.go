package category_manage_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type CategoryManageDao struct {
	Orm *gorm.DB
}

func NewCategoryManageDao() CategoryManageDao {
	return CategoryManageDao{
		Orm: global.MDB,
	}
}
