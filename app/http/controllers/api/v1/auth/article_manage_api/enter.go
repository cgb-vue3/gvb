package article_manage_api

import (
	"akita/app/http/servers/article_manage_server"
)

type Api struct {
	articleManageServers *article_manage_server.ArticleServers // 图片
}

func NewArticleManageApi() *Api {
	return &Api{
		articleManageServers: article_manage_server.NewArticleServers(),
	}
}
