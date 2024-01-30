package category_manage_api

import (
	"akita/app/http/servers/category_manage_server"
)

type Api struct {
	categoryManageServers *category_manage_server.CategoryServers
}

func NewArticleManageApi() *Api {
	return &Api{
		categoryManageServers: category_manage_server.NewCategoryServers(),
	}
}
