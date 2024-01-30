package side_menu_api

import (
	"akita/app/http/servers/side_menu_server"
)

type Api struct {
	sideMenuServers *side_menu_server.SideMenuServers // 图片
}

func NewUserManageApi() *Api {
	return &Api{
		sideMenuServers: side_menu_server.NewSideMenuServers(),
	}
}
