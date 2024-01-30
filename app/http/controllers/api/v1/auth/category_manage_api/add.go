package category_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/category_manage_api/category_manage_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) Add(ctx *gin.Context) {
	var addParams category_manage_params.AddParams
	err := ctx.ShouldBindJSON(&addParams)
	if err != nil {
		global.Mlog.Error("添加文章参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.categoryManageServers.Add(ctx, addParams)
}
