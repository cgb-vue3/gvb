package image_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type ImageDao struct {
	Orm *gorm.DB
}

func NewImgDao() ImageDao {
	return ImageDao{
		Orm: global.MDB,
	}
}
