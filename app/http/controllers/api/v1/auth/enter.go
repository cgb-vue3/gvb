package auth

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api"
	"akita/app/http/controllers/api/v1/auth/category_manage_api"
	"akita/app/http/controllers/api/v1/auth/child_menu_api"
	"akita/app/http/controllers/api/v1/auth/comment_manage_api"
	"akita/app/http/controllers/api/v1/auth/side_menu_api"
	"akita/app/http/controllers/api/v1/auth/tag_manage_api"
	"akita/app/http/controllers/api/v1/auth/user_manage_api"
)

type AuthGroupApis struct {
	UserManage     *user_manage_api.Api
	SideMenu       *side_menu_api.Api
	ChildMenu      *child_menu_api.Api
	ArticleManage  *article_manage_api.Api
	CategoryManage *category_manage_api.Api
	TagManage      *tag_manage_api.Api
	CommentManage  *comment_manage_api.Api
}

func NewAuthApis() *AuthGroupApis {
	return &AuthGroupApis{
		user_manage_api.NewUserManageApi(),
		side_menu_api.NewUserManageApi(),
		child_menu_api.NewChildMenuApi(),
		article_manage_api.NewArticleManageApi(),
		category_manage_api.NewArticleManageApi(),
		tag_manage_api.NewArticleManageApi(),
		comment_manage_api.NewCommentManageApi(),
	}
}
