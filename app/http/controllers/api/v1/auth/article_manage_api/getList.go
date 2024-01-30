package article_manage_api

import (
	"akita/app/http/controllers/common"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) GetList(ctx *gin.Context) {
	var pagParams common.PagingParams

	err := ctx.ShouldBindQuery(&pagParams)
	if err != nil {
		global.Mlog.Error("返回文章列表参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.articleManageServers.GetList(ctx, pagParams)
}
