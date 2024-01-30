package user_api

import (
	"akita/app/http/servers/user_server"
)

type Api struct {
	userServers *user_server.UserServers
}

func NewUserApi() *Api {
	return &Api{
		userServers: user_server.NewUserServers(),
	}
}
