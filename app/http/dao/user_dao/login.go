package user_dao

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/encryption"
	"akita/pkg/jwt"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M UserDao) Login(ctx *gin.Context, loginParams user_params.LoginParams) {
	var loginModel models.UserModel
	// 通过用户名判断用户是否存在
	if err := M.Orm.Where("user_name = ?", loginParams.UserName).Take(&loginModel).Error; err != nil {
		// 用户不存在
		global.Mlog.Error("用户不存在", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserNoExist))
		return
	}
	// 用户存在，对比密码
	pwd := []byte(loginParams.PassWord)
	pwdVerify := encryption.ComparePasswords(loginModel.PassWord, pwd)
	if !pwdVerify {
		// 用户不存在
		global.Mlog.Warn("密码错误")
		response.Err400(ctx, response.CodeWithMsg(response.CodePasswordFailed))
		return
	}

	// 生成token
	token, err := jwt.GenToken(loginModel.ID, loginModel.Nickname, int(loginModel.Role))
	if err != nil {
		global.Mlog.Error("token生成失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeLoginSuccess))
		return
	}
	// 登录成功,返回token和用户信息
	response.OK200(ctx, response.CodeWithMsg(response.CodeLoginSuccess), response.WithData(map[string]any{
		"token": token,
	}))
}
