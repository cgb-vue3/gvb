package user_manage_dao

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"akita/app/http/models"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
)

func (M UserManageDao) Del(ctx *gin.Context, delParam user_manage_params.DelParams) {
	var userModel []models.UserModel
	// 通过del参数查找图片
	M.Orm.Where("id in ?", delParam.IDList).Delete(&userModel)
	response.OK200(ctx,
		response.CodeWithMsg(response.CodeDelSucceed))
}
