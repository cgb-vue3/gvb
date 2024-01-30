package user_manage_server

import (
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) GetTotal(ctx *gin.Context) {
	M.Dao.GetTotal(ctx)
}
