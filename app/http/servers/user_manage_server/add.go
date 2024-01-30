package user_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) Add(ctx *gin.Context, addParam user_manage_params.AddParams) {
	M.Dao.UserManageDao.Add(ctx, addParam)
}
