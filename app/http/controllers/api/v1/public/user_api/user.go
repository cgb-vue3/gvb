package user_api

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func (M Api) Register(ctx *gin.Context) {
	var registerParams user_params.RegisterParams
	err := ctx.ShouldBindJSON(&registerParams)
	if err != nil {
		global.Mlog.Error("参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userServers.Register(ctx, registerParams)
}

// Login 用户登录
func (M Api) Login(ctx *gin.Context) {
	var loginParams user_params.LoginParams
	err := ctx.ShouldBindJSON(&loginParams)
	if err != nil {
		global.Mlog.Error("参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userServers.Login(ctx, loginParams)
}

// SendEmail  发送邮件
func (M Api) SendEmail(ctx *gin.Context) {
	var (
		emailParams user_params.EmailParams
		//respUserInfo user_resp.RespUserInfo
	)
	err := ctx.ShouldBindJSON(&emailParams)
	if err != nil {
		global.Mlog.Error("参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userServers.SendEmail(ctx, emailParams)
}

// ChangePwd 修改密码
func (M Api) ChangePwd(ctx *gin.Context) {
	var changePwdParams user_params.ChangePwdParams
	err := ctx.ShouldBindJSON(&changePwdParams)
	if err != nil {
		global.Mlog.Error("参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userServers.ChangePwd(ctx, changePwdParams)
}
