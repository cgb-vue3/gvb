package user_server

import (
	"akita/app/http/controllers/api/v1/public/user_api/user_params"
	"akita/global"
	"akita/pkg/redis"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

func (M UserServers) ChangePwd(ctx *gin.Context, changePwdParams user_params.ChangePwdParams) {
	// 判断验证码是否存在正确。
	// 获取redis中的验证码
	get := redis.Get(changePwdParams.Code)
	// 判断redis是否为空，且id值与用户传的id是否相等
	if get != "" && cast.ToUint(get) == changePwdParams.ID {
		// 判断用户两次输入的密码是否一致
		if changePwdParams.PassWord == changePwdParams.RePassWord {
			// 清空掉redis中的验证码
			err := redis.DelStr(changePwdParams.Code)
			if err != nil {
				global.Mlog.Error("删除验证码错误", zap.Error(err))
				return
			}
			M.Dao.UserDao.ChangePwd(ctx, changePwdParams)
		}
		return
	}
	global.Mlog.Error("验证码错误或已被使用,请重新发送")
	response.Err400(ctx, response.CodeWithMsg(response.CodeCaptchaError))
}
