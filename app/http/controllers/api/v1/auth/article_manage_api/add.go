package article_manage_api

import (
	"akita/app/http/controllers/api/v1/auth/article_manage_api/article_manage_params"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
)

//	@Tags		文章管理
//	@Summary	添加文章
//	@Accept		multipart/form-data
//	@Produce	json
//	@Param		data	body		article_manage_params.AddParams	true	"添加图片的参数"
//	@Success	200		{object}	response.go.JSON{}				"上传成功"
//	@Failure	400		{object}	response.go.JSON{}				"上传失败"
//	@Router		/api/auth/v1/article/add [post]
//
// Add 添加文章
func (M Api) Add(ctx *gin.Context) {
	var addParams article_manage_params.AddParams
	err := ctx.ShouldBindJSON(&addParams)
	if err != nil {
		global.Mlog.Error("添加文章参数错误")
		valida.Validator(ctx, err)
		return
	}
	M.articleManageServers.Add(ctx, addParams)
}
