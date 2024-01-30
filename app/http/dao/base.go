package dao

import (
	"akita/app/http/dao/article_manage_dao"
	"akita/app/http/dao/category_manage_dao"
	"akita/app/http/dao/child_menu_dao"
	"akita/app/http/dao/comment_manage_dao"
	"akita/app/http/dao/image_dao"
	"akita/app/http/dao/menu_dao"
	"akita/app/http/dao/side_menu_dao"
	"akita/app/http/dao/tag_manage_dao"
	"akita/app/http/dao/user_dao"
	"akita/app/http/dao/user_manage_dao"
	"sync"
)

type BaseDao struct {
	image_dao.ImageDao
	menu_dao.MenuDao
	user_dao.UserDao
	user_manage_dao.UserManageDao
	side_menu_dao.SideMenuDao
	child_menu_dao.ChildMenuDao
	article_manage_dao.ArticleManageDao
	category_manage_dao.CategoryManageDao
	tag_manage_dao.TagManageDao
	comment_manage_dao.CommentManageDao
}

var (
	baseDao *BaseDao
	once    sync.Once
)

// NewBaseDao 实例化BaseDao函数
func NewBaseDao() *BaseDao {
	once.Do(func() {
		if baseDao == nil {
			baseDao = &BaseDao{
				image_dao.NewImgDao(),
				menu_dao.NewMenuDao(),
				user_dao.NewUserDao(),
				user_manage_dao.NewUserManageDao(),
				side_menu_dao.NewSideMenuDao(),
				child_menu_dao.NewChildMenuDao(),
				article_manage_dao.NewArticleManageDao(),
				category_manage_dao.NewCategoryManageDao(),
				tag_manage_dao.NewTagManageDao(),
				comment_manage_dao.NewCommentManageDao(),
			}
		}
	})

	return baseDao
}
