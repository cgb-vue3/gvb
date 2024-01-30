package comment_manage_dao

import (
	"akita/global"
	"gorm.io/gorm"
)

type CommentManageDao struct {
	Orm *gorm.DB
}

func NewCommentManageDao() CommentManageDao {
	return CommentManageDao{
		Orm: global.MDB,
	}
}
