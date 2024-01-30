package article_manage_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type ArticleManageDao struct {
	Orm *gorm.DB
}

func NewArticleManageDao() ArticleManageDao {
	return ArticleManageDao{
		Orm: global.MDB,
	}
}
