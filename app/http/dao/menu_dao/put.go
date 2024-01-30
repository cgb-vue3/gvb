package menu_dao

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Put 编辑菜单
func (M MenuDao) Put(ctx *gin.Context, putMenuParam menu_params.PutMenuParams) {
	var (
		menuModelList []models.MenuModel
		menuModel     models.MenuModel
		imageModel    []models.ImageModel
		imageIDList   []uint
	)

	// 查询到所有菜单数据
	M.Orm.Find(&menuModelList)
	// 判断修改后的菜单名是否重复
	// 循环判断菜单id是否与传入菜单id的值是否相等，如果相等跳出循环，
	//如果不相等就去判断传入的title在菜单表中是否存在，如果存在返回菜单名已存在
	for _, model := range menuModelList {
		if model.ID == putMenuParam.ID {
			continue
		}
		if err := M.Orm.Take(&model, "title = ?", putMenuParam.Title).Error; err != nil {
			// 菜单名没有重复
			continue
		}
		// 菜单名重复
		response.Err400(ctx, response.WithMsg("菜单名已存在，请重新输入"))
		return
	}

	// 查到当前传入的菜单
	if err := M.Orm.Where("id = ?", putMenuParam.ID).Find(&menuModel).Error; err != nil {
		response.Err400(ctx, response.WithMsg("没有找到该菜单，更新失败"))
		return
	}

	// 清除关联
	if err := M.Orm.Model(&menuModel).Association("ImageModel").Clear(); err != nil {
		global.Mlog.Error("清空关联失败", zap.Error(err))
		response.Err400(ctx, response.WithMsg("清空关联失败,更新失败"))
		return
	}

	// 入库前判断有没有传图片id，没有直接跟新，有就重新关联图片
	if len(putMenuParam.ImageList) == 0 {
		var putModel = models.MenuModel{
			Title:        putMenuParam.Title,
			TitleEn:      putMenuParam.TitleEn,
			Slogan:       putMenuParam.Slogan,
			ImageCutTime: putMenuParam.ImageCutTime,
			Sort:         putMenuParam.Sort,
		}
		err := M.Orm.Model(&menuModel).Updates(&putModel).Error
		if err != nil {
			global.Mlog.Error("菜单更新失败", zap.Error(err))
			response.Err400(ctx, response.WithMsg("菜单更新失败"))
			return
		}
	}

	// 如果传入了图片id，先判断图片是否存在
	for _, imageList := range putMenuParam.ImageList {
		// 先判断图片是否存在，存在便去添加关联
		if err := M.Orm.Take(&imageModel, "id = ?", imageList.ID).Error; err != nil {
			global.Mlog.Warn("图片不存在")
			continue
		}
		imageIDList = append(imageIDList, imageList.ID)
	}

	// 将图片id和菜单id添加到关联表中
	for _, imgID := range imageIDList {
		menuImageModel := models.MenuImagesModel{
			MenuModelID:  menuModel.ID,
			ImageModelID: imgID,
		}
		// 更新关联表
		if err := M.Orm.Create(&menuImageModel).Error; err != nil {
			global.Mlog.Error("菜单更新失败", zap.Error(err))
			response.Err400(ctx, response.WithMsg("菜单更新失败"))
			return
		}
	}

	response.OK200(ctx, response.WithMsg("菜单更新成功"))

}
