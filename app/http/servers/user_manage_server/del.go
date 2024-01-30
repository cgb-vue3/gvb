package user_manage_server

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"github.com/gin-gonic/gin"
)

func (M UserManageServers) Del(ctx *gin.Context, delParam user_manage_params.DelParams) {
	M.Dao.UserManageDao.Del(ctx, delParam)
}
