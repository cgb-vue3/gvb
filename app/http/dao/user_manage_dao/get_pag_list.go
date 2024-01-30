package user_manage_dao

import (
	"akita/app/http/controllers/common"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/paging"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
)

func (M UserManageDao) UserPagList(ctx *gin.Context, id uint, role int, pagParams common.PagingParams) {
	var userModel []models.UserModel
	// 获取全部用户和当前用户的用户级别
	M.Orm.Find(&userModel).Where("id = ?", id)

	// 用户等级大于等于4返回权限不足

	if role >= int(models.PermissionVisitor) {
		global.Mlog.Error("用户权限不足，返回用完列表失败")
		response.Err401(ctx)
		return
	}

	// 按用级别返回用户信息
	cutList := userLevel(userModel, role, id)

	userList, err := paging.Pag(
		cutList,
		M.Orm,
		paging.Option{
			Params: pagParams,
		}, "find")
	if err != nil {
		response.Err401(
			ctx,
			response.CodeWithMsg(response.CodeUserListFailed),
			response.WithErr(err),
		)
		return
	}

	//// 排序、分页参数
	//sort, offset := paging.PagWithSort(pagParams)
	//err := M.Orm.Limit(pagParams.PageSize).Offset(offset).Order(sort).Find(&cutList)

	response.OK200(ctx,
		response.CodeWithMsg(response.CodeUserListSucceed),
		response.WithData(gin.H{
			"total":    len(cutList),
			"listData": userList,
		}))
}

func userLevel(userList []models.UserModel, role int, id uint) []models.UserModel {
	// role: 1 超级管理员 2 管理员 3 普通用户 4 游客 5 被禁用的用户
	userList = cutUser(userList, role, id)
	return userList
}

func cutUser(userList []models.UserModel, role int, id uint) []models.UserModel {
	//将当前用户截取出去
	cut := make([]models.UserModel, 0)
	for _, model := range userList {
		//// 用户邮箱脱敏
		//model.Email = data_mask.EmailMasking(model.Email)
		//// 用户手机号脱敏
		//model.Phone = data_mask.PhoneMasking(model.Phone)
		if int(model.Role) >= role {
			switch role {
			//	超级管理员
			case int(models.PermissionSuperAdmin):
				cut = append(cut, model)
				// 管理员
			case int(models.PermissionAdin):
				cut = append(cut, model)
			}
		}
	}
	return cut
}
