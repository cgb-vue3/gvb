package user_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/encryption"
	"akita/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func (M UserManageDao) Add(ctx *gin.Context, addParam user_manage_params.AddParams) {

	fmt.Println(addParam)
	// 判断用户名是否存在，存在返回错误信息，不存在就创建用户
	if err := M.Orm.Where("user_name = ?", addParam.UserName).Take(&models.UserModel{}).Error; err == nil {
		global.Mlog.Error("用户已存在，添加失败")
		response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodeUserExist))
		return
	}

	// 判断手机号是否已被注册
	if err := M.Orm.Where("phone = ?", addParam.Phone).Take(&models.UserModel{}).Error; err == nil {
		// 返回用户存在信息
		global.Mlog.Error("手机号已被注册，添加失败")
		response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodePhoneExist))
		return
	}
	// 判断邮箱是否已被注册
	if err := M.Orm.Where("email = ?", addParam.Email).Take(&models.UserModel{}).Error; err == nil {
		global.Mlog.Error("邮箱已被注册，添加失败")
		response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodeEmailExist))
		return
	}
	// 用户密码加密
	pwd := []byte(addParam.PassWord)
	hashPwd, err := encryption.HashAndSalt(pwd)
	if err != nil {
		global.Mlog.Error("用户创建失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserAddFailed))
		return
	}

	createUserModel := &models.UserModel{
		Nickname: addParam.NickName,
		UserName: addParam.UserName,
		Phone:    addParam.Phone,
		Role:     models.Role(cast.ToInt32(addParam.Role)),
		Avatar:   addParam.Avatar,
		PassWord: hashPwd,
		Email:    addParam.Email,
	}

	// 创建用户
	err = M.Orm.Create(&createUserModel).Error
	if err != nil {
		global.Mlog.Error("用户创建失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserAddFailed))
		return
	}
	global.Mlog.Info("用户创建成功")
	response.OK201(ctx, response.CodeWithMsg(response.CodeUserAddSuccess))

}
