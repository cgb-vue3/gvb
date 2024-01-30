package comment_manage_api

import (
	"akita/app/http/servers/comment_manage_server"
)

type Api struct {
	commentManageServers *comment_manage_server.CommentServers
}

func NewCommentManageApi() *Api {
	return &Api{
		commentManageServers: comment_manage_server.NewCommentServers(),
	}
}
