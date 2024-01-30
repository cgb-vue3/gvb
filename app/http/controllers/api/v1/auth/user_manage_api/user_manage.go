package user_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/user_manage_api/user_manage_params"
	"akita/app/http/controllers/common"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

// UserPagList 返回用户列表
func (M Api) UserPagList(ctx *gin.Context) {
	var pagParams common.PagingParams
	err := ctx.ShouldBindQuery(&pagParams)
	if err != nil {
		global.Mlog.Error("用户分页参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userManageServers.UserPagList(ctx, pagParams)
}

// GetInfo 返回当前用户信息
func (M Api) GetInfo(ctx *gin.Context) {
	M.userManageServers.GetInfo(ctx)
}

// Add 添加用户
func (M Api) Add(ctx *gin.Context) {
	var addParam user_manage_params.AddParams
	err := ctx.ShouldBindJSON(&addParam)
	if err != nil {
		global.Mlog.Error("添加用户参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userManageServers.Add(ctx, addParam)
}

// Del 删除用户
func (M Api) Del(ctx *gin.Context) {
	var delParam user_manage_params.DelParams
	err := ctx.ShouldBindJSON(&delParam)
	if err != nil {
		global.Mlog.Error("删除用户参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userManageServers.Del(ctx, delParam)
}

// Put 编辑用户
func (M Api) Put(ctx *gin.Context) {
	var putParams user_manage_params.PutParams
	err := ctx.ShouldBindJSON(&putParams)
	if err != nil {
		global.Mlog.Error("编辑用户参数绑定错误")
		valida.Validator(ctx, err)
		return
	}
	M.userManageServers.Put(ctx, putParams)
}

// 获取用户文章、用户数量

func (M Api) GetTotal(ctx *gin.Context) {
	M.userManageServers.GetTotal(ctx)
}
