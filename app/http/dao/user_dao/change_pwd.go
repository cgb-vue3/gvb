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

func (M UserDao) ChangePwd(ctx *gin.Context, changePwdParams user_params.ChangePwdParams) {
	var (
		userModel models.UserModel
		id        = changePwdParams.ID
		passWord  = changePwdParams.PassWord
	)
	// 通过id查找用户
	if err := M.Orm.Where("id=?", id).Take(&userModel).Error; err != nil {
		global.Mlog.Error("用户不存在", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeChangePwdFailed))
		return
	}
	pwdHash, _ := encryption.HashAndSalt([]byte(passWord))
	// 修改密码
	if err := M.Orm.Model(&userModel).Select("pass_word").Updates(models.UserModel{PassWord: pwdHash}).Error; err != nil {
		global.Mlog.Error("密码修改失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeChangePwdFailed))
		return
	}
	response.OK200(ctx, response.CodeWithMsg(response.CodeChangePwdSucceed))
}
