package image_server

import (
	"akita/app/http/controllers/common"
	"github.com/gin-gonic/gin"
)

// ResponsePagingList 返回图片分页
func (M ImageServers) ResponsePagingList(ctx *gin.Context, pagParams common.PagingParams) {
	M.Dao.ImageDao.ResponsePagingList(ctx, pagParams)
}
