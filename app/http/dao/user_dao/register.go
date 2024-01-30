package user_dao

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/encryption"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M UserDao) Register(ctx *gin.Context, registerParams user_params.RegisterParams) {

	// 判断用户名是否存在，存在返回错误信息，不存在就创建用户
	if err := M.Orm.Where("user_name = ?", registerParams.UserName).Take(&models.UserModel{}).Error; err == nil {
		// 返回用户存在信息
		global.Mlog.Error("用户名已存在，注册失败")
		response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodeUserExist))
		return
	}
	if err := M.Orm.Where("email = ?", registerParams.Email).Take(&models.UserModel{}).Error; err == nil {
		// 返回用户存在信息
		global.Mlog.Error("邮箱已存在，注册失败")
		response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodeEmailExist))
		return
	}

	// 用户密码加密
	pwd := []byte(registerParams.PassWord)
	hashPwd, err := encryption.HashAndSalt(pwd)
	if err != nil {
		global.Mlog.Error("用户创建失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserCreationFailed))
		return
	}

	createUserModel := &models.UserModel{
		Nickname: registerParams.NickName,
		UserName: registerParams.UserName,
		PassWord: hashPwd,
		Email:    registerParams.Email,
	}
	// 创建用户
	err = M.Orm.Create(&createUserModel).Error
	if err != nil {
		global.Mlog.Error("用户创建失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserCreationFailed))
		return
	}
	global.Mlog.Info("用户创建成功")
	response.OK201(ctx, response.CodeWithMsg(response.CodeUserCreatedSuccess))

}
