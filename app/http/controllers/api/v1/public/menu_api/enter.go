package menu_api

import (
	"akita/app/http/servers/menu_Server"
)

type Api struct {
	menuServer *menu_server.MenuServers // 图片
}

func NewMenuApi() *Api {
	return &Api{
		menuServer: menu_server.NewMenuServer(),
	}
}
