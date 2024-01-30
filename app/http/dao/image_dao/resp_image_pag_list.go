package image_dao

import (
	"akita/app/http/controllers/common"
	"akita/app/http/models"
	"akita/pkg/paging"
	"akita/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ResponsePagingList 返回图片分页列表
func (M ImageDao) ResponsePagingList(ctx *gin.Context, pagParams common.PagingParams) {
	var (
		imageModel []models.ImageModel
		count      int64
	)

	M.Orm.Find(&imageModel).Count(&count)
	pag, err := paging.Pag(imageModel, M.Orm, paging.Option{Params: pagParams}, "find")
	if err != nil {
		fmt.Println(err)
		return
	}
	imageModel = pag

	// 成功的响应
	response.OK200(ctx,
		response.WithCode(response.CodeImageRespSucceed),
		response.WithData(map[string]any{
			"total":    count,
			"listData": imageModel,
		}))
}
