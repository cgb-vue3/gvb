package comment_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/comment_manage_api/comment_manage_params"
	"github.com/gin-gonic/gin"
)

func (M CommentServers) Add(ctx *gin.Context, commentParams comment_manage_params.CommentParams) {
	M.Dao.CommentManageDao.Add(ctx, commentParams)
}
