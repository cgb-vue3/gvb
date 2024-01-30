package article_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) Del(ctx *gin.Context) {
	var delParams article_manage_params.DelParams
	err := ctx.ShouldBindJSON(&delParams)
	if err != nil {
		global.Mlog.Error("删除文章参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.articleManageServers.Del(ctx, delParams)
}
