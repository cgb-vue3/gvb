package tag_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/tag_manage_api/tag_manage_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) Add(ctx *gin.Context) {
	var addParams tag_manage_params.AddParams
	err := ctx.ShouldBindJSON(&addParams)
	if err != nil {
		global.Mlog.Error("添加文章参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.tagManageServers.Add(ctx, addParams)
}
