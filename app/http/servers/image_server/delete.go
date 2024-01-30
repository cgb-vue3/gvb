package image_server

import (
	"akita/app/http/controllers/api/v1/public/image_api/image_params"
	"akita/app/http/controllers/api/v1/public/image_api/image_resp"
	"github.com/gin-gonic/gin"
)

// Delete 删除图片
func (M ImageServers) Delete(ctx *gin.Context, deleteParams image_params.DeleteImgParams, respDeleteInfoList []image_resp.RespDeleteInfoList) {
	M.Dao.ImageDao.Delete(ctx, deleteParams, respDeleteInfoList)
}
