package image_dao

import (
	"akita/app/http/controllers/api/v1/public/image_api/image_params"
	"akita/app/http/controllers/api/v1/public/image_api/image_resp"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"akita/pkg/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

// Delete 删除图片
func (M ImageDao) Delete(ctx *gin.Context, deleteParams image_params.DeleteImgParams, respDeleteInfoList []image_resp.RespDeleteInfoList) {
	var (
		total      int
		imageModel models.ImageModel
	)
	// 通过id拿到图片集合。
	for _, id := range deleteParams.IDList {
		row := M.Orm.Where("id = ?", id).Find(&imageModel).RowsAffected
		if row == 0 {
			respDeleteInfoList = responseList(respDeleteInfoList, id, imageModel.Name, "图片不存在")
			continue
		}
		if err := M.Orm.Delete(&imageModel).Error; err != nil {
			global.Mlog.Error("图片删除失败", zap.Error(err))
			respDeleteInfoList = responseList(respDeleteInfoList, imageModel.ID, imageModel.Name, "图片删除成功")
			continue
		}

		if imageModel.Env == 0 {
			err := os.RemoveAll(imageModel.Path)
			if err != nil {
				global.Mlog.Error("删除本地图片失败", zap.Error(err))
				respDeleteInfoList = responseList(respDeleteInfoList, imageModel.ID, imageModel.Name, "删除本地图片失败")
				continue
			}
		} else if imageModel.Env == 1 {
			// 七牛批量删除图片
			upload.DeleteMassQiNiuImages(imageModel.Name)
		}
		total++
		respDeleteInfoList = responseList(respDeleteInfoList, imageModel.ID, imageModel.Name, "图片删除成功")

	}
	response.OK200(ctx, response.WithData(map[string]any{
		"total":       total,
		"delete_list": respDeleteInfoList,
	}))
}

func responseList(respDeleteInfoList []image_resp.RespDeleteInfoList, id uint, name, msg string) []image_resp.RespDeleteInfoList {
	return append(respDeleteInfoList, image_resp.RespDeleteInfoList{
		ID:   id,
		Name: name,
		Msg:  msg,
	})
}
