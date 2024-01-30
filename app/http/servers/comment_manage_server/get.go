package comment_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/comment_manage_api/comment_manage_params"
	"github.com/gin-gonic/gin"
)

func (M CommentServers) Get(ctx *gin.Context, getParams comment_manage_params.Get) {
	//M.Dao.Get(ctx, getParams)
}
