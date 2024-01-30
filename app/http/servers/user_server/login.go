package user_server

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"github.com/gin-gonic/gin"
)

func (M UserServers) Login(ctx *gin.Context, loginParams user_params.LoginParams) {
	M.Dao.UserDao.Login(ctx, loginParams)
}
