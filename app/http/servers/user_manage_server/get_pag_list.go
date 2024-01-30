package user_manage_server

import (
	"akita/app/http/controllers/common"
	"akita/pkg/get_claims"
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) UserPagList(ctx *gin.Context, pagParams common.PagingParams) {
	// 获取用户信息
	id, role, _ := get_claims.GetClaims(ctx)
	M.Dao.UserManageDao.UserPagList(ctx, id, role, pagParams)
}
