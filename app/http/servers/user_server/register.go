package user_server

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
)

func (M UserServers) Register(ctx *gin.Context, registerParams user_params.RegisterParams) {
	// 对参数进行校验
	// 判断password和re_password是否相等
	if registerParams.PassWord != registerParams.RePassWord {
		response.Err422(
			ctx,
			response.CodeWithMsg(response.CodeInequalityPassword))
		return
	}
	M.Dao.UserDao.Register(ctx, registerParams)
}
