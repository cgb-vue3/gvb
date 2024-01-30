package user_manage_server

import (
	"akita/pkg/get_claims"
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) GetInfo(ctx *gin.Context) {
	id, _, _ := get_claims.GetClaims(ctx)
	M.Dao.UserManageDao.GetInfo(ctx, id)
}
