package side_menu_server

import (
	"akita/app/http/controllers/common"
	"akita/pkg/get_claims"
	"github.com/gin-gonic/gin"
)

func (M SideMenuServers) GetPag(ctx *gin.Context, sidePagParams common.PagingParams) {
	_, role, _ := get_claims.GetClaims(ctx)
	M.Dao.SideMenuDao.GetPag(ctx, sidePagParams, role)
}
