package child_menu_api

import (
	"akita/app/http/servers/child_menu_server"
)

type Api struct {
	childMenuServers *child_menu_server.ChildMenuServers // 图片
}

func NewChildMenuApi() *Api {
	return &Api{
		childMenuServers: child_menu_server.NewChildMenuServers(),
	}
}
