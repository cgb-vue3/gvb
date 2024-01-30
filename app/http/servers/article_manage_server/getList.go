package article_manage_server

import (
	"akita/app/http/controllers/common"
	"akita/pkg/get_claims"
	"github.com/gin-gonic/gin"
)

func (M ArticleServers) GetList(ctx *gin.Context, pagParams common.PagingParams) {
	// 获取用户信息
	id, _, _ := get_claims.GetClaims(ctx)
	M.Dao.ArticleManageDao.GetList(ctx, pagParams, id)
}
