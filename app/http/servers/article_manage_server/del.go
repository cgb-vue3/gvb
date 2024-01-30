package article_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"github.com/gin-gonic/gin"
)

func (M ArticleServers) Del(ctx *gin.Context, delParams article_manage_params.DelParams) {
	M.Dao.ArticleManageDao.Del(ctx, delParams)
}
