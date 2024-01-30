package article_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"github.com/gin-gonic/gin"
)

func (M ArticleServers) Add(ctx *gin.Context, addParams article_manage_params.AddParams) {
	M.Dao.ArticleManageDao.Add(ctx, addParams)
}
