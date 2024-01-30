package user_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/encryption"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func (M UserManageDao) Put(ctx *gin.Context, putParams user_manage_params.PutParams) {
	var putModel models.UserModel
	var userModel models.UserModel
	// 根据id查到用户
	if err := M.Orm.Where("id = ?", putParams.ID).Take(&putModel).Error; err != nil {
		global.Mlog.Error("用户不存在", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserNoExist))
		return
	}

	// 判断用户是要编辑，还是重置密码
	if putParams.Type == "rePassword" {
		// 用户密码加密
		pwd := []byte(putParams.PassWord)
		hashPwd, err := encryption.HashAndSalt(pwd)
		if err != nil {
			global.Mlog.Error("密码重置失败", zap.Error(err))
			response.Err400(ctx, response.CodeWithMsg(response.CodeReSetPwdFailed))
			return
		}

		// 重置密码
		if err := M.Orm.Model(&putModel).Select("PassWord").
			Updates(models.UserModel{
				PassWord: hashPwd,
			}).
			Error; err != nil {
			global.Mlog.Error("密码重置失败", zap.Error(err))
			response.Err400(ctx, response.CodeWithMsg(response.CodeReSetPwdFailed))
			return
		}
		global.Mlog.Info("密码重置成功")
		response.OK200(ctx, response.CodeWithMsg(response.CodeReSetPwdSuccess), response.WithData(map[string]any{
			"data": putParams,
		}))
		return
	}

	// 根据手机号查找该用户
	// 将查到的id与传入的id对比，如果id相等就允许修改
	// 如果不等就返回手机号已存在
	// 判断手机号是否已被注册
	if err := M.Orm.Where("phone = ?", putParams.Phone).Take(&userModel).Error; err == nil {
		if userModel.ID != putParams.ID {
			// 返回用户存在信息
			global.Mlog.Error("手机号已被注册，编辑失败")
			response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodePhoneExist))
			return
		}
	}
	// 判断邮箱是否已被注册
	if err := M.Orm.Where("email = ?", putParams.Email).Take(&userModel).Error; err == nil {
		if userModel.ID != putParams.ID {
			global.Mlog.Error("邮箱已被注册，编辑失败")
			response.Err400(ctx, response.WithCode(response.CodeFailed), response.CodeWithMsg(response.CodeEmailExist))
			return
		}
	}

	// 更新
	if err := M.Orm.Model(&putModel).Select("nickname", "email", "phone", "role", "avatar").
		Updates(models.UserModel{
			Nickname: putParams.NickName,
			Email:    putParams.Email,
			Phone:    putParams.Phone,
			Role:     models.Role(cast.ToInt32(putParams.Role)),
			Avatar:   putParams.Avatar,
		}).
		Error; err != nil {
		global.Mlog.Error("用户更新失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeUserUpdateFailed))
		return
	}
	global.Mlog.Info("用户更新成功")
	response.OK200(ctx, response.CodeWithMsg(response.CodeUserUpdateSuccess))
}
