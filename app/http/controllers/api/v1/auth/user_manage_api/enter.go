package user_manage_api

import (
	"akita/app/http/servers/user_manage_server"
)

type Api struct {
	userManageServers *user_manage_server.UserManageServers // 图片
}

func NewUserManageApi() *Api {
	return &Api{
		userManageServers: user_manage_server.NewUserManageServers(),
	}
}
