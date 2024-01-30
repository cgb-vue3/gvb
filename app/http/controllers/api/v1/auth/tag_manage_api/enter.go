package tag_manage_api

import (
	"akita/app/http/servers/tag_manage_server"
)

type Api struct {
	tagManageServers *tag_manage_server.TagServers
}

func NewArticleManageApi() *Api {
	return &Api{
		tagManageServers: tag_manage_server.NewTagServers(),
	}
}
