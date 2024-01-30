package user_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) Put(ctx *gin.Context, putParams user_manage_params.PutParams) {
	M.Dao.UserManageDao.Put(ctx, putParams)
}
