package tag_manage_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type TagManageDao struct {
	Orm *gorm.DB
}

func NewTagManageDao() TagManageDao {
	return TagManageDao{
		Orm: global.MDB,
	}
}
