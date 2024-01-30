package image_api

import (
	"akita/app/http/controllers/api/v1/public/image_api/image_params"
	"akita/app/http/controllers/api/v1/public/image_api/image_resp"
	"akita/app/http/controllers/common"
	"akita/global"
	"akita/pkg/valida"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//	@Tags		图片管理
//	@Summary	上传图片
//	@Accept		json
//	@Produce	json
//	@Param		data	body		dto.AddImgParams							true	"添加图片的参数"
//	@Success	200		{object}	response.go.JSON{data=[]dto.AddImgParams}	"上传成功"
//	@Failure	400		{object}	response.go.JSON{}							"上传失败"
//	@Router		/api/v1/images [post]
//
// Add 上传图片api
func (M Api) Add(ctx *gin.Context) {

	var respAddImageInfoList = make([]image_resp.RespAddImageInfoList, 0)
	M.imageServer.Add(ctx, respAddImageInfoList)
}

//	@Tags		图片管理
//	@Summary	返回图片列表
//	@Produce	json
//	@Param		data	query		image_params.FindPagingParams							true	"返回页数"
//	@Success	200		{object}	response.go.JSON{data=[]image_params.FindPagingParams}	"成功"
//	@Failure	400		{object}	response.go.JSON{data=string}							"失败"
//	@Router		/api/v1/images [get]
//
// ResponsePagingList 分页查询图片，返回图片列表
func (M Api) ResponsePagingList(ctx *gin.Context) {
	var pagParams common.PagingParams

	err := ctx.ShouldBindQuery(&pagParams)
	if err != nil {
		global.Mlog.Error("图片分页参数绑定失败", zap.Error(err))
		valida.Validator(ctx, err)
		return
	}
	M.imageServer.ResponsePagingList(ctx, pagParams)
}

//	@Tags		图片管理
//	@Summary	图片删除
//	@Produce	json
//	@Param		IDList	body		dto.RemoveDTO		true	"图片id切片"
//	@Success	200		{object}	response.go.JSON{}	"删除成功"
//	@Failure	400		{object}	response.go.JSON{}	"图片未找到或已删除"
//	@Router		/api/v1/images [delete]
//
// Delete 删除数据库图片和本地图片
func (M Api) Delete(ctx *gin.Context) {
	var (
		deleteParams       image_params.DeleteImgParams
		respDeleteInfoList = make([]image_resp.RespDeleteInfoList, 0)
	)

	err := ctx.ShouldBindJSON(&deleteParams)
	if err != nil {
		global.Mlog.Error("图片分页参数绑定失败", zap.Error(err))
		valida.Validator(ctx, err)
		return
	}

	M.imageServer.Delete(ctx, deleteParams, respDeleteInfoList)
}
