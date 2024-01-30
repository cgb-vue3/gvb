package tag_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/tag_manage_api/tag_manage_params"
	"github.com/gin-gonic/gin"
)

func (M TagServers) Add(ctx *gin.Context, addParams tag_manage_params.AddParams) {
	M.Dao.TagManageDao.Add(ctx, addParams)
}
