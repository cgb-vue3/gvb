package comment_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/comment_manage_api/comment_manage_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

func (M Api) Get(ctx *gin.Context) {
	var getParams comment_manage_params.Get
	err := ctx.ShouldBindQuery(&getParams)
	if err != nil {
		global.Mlog.Error("添加评论参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.commentManageServers.Get(ctx, getParams)
}
