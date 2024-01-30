package menu_dao

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_params"
	model "akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Add 添加菜单
func (M MenuDao) Add(ctx *gin.Context, menuAddParams menu_params.AddMenuParams) {
	var (
		menuImgModel []model.MenuImagesModel
		menuModel    = &model.MenuModel{
			Title:        menuAddParams.Title,
			TitleEn:      menuAddParams.TitleEn,
			Slogan:       menuAddParams.Slogan,
			ImageCutTime: menuAddParams.ImageCutTime,
			Sort:         menuAddParams.Sort,
		}
	)

	// 入库前先判断该字段
	affected := M.Orm.Where("title = ?", menuModel.Title).Find(&model.MenuModel{}).RowsAffected
	fmt.Println(affected)
	if affected > 0 {
		global.Mlog.Error("菜单已存在")
		response.Err400(ctx, response.WithMsg("菜单已存在,添加失败"))
		return
	}
	//创建menu数据入库
	err := M.Orm.Create(&menuModel).Error
	if err != nil {
		global.Mlog.Error("菜单创建失败:", zap.Error(err))
		response.Err400(ctx, response.WithMsg("菜单创建失败"))
		return
	}

	// 判断是否传了图片id，没有直接返回创建成功
	if len(menuAddParams.ImageList) == 0 {
		response.OK201(
			ctx,
			response.WithMsg("菜单创建成功"),
			response.WithData(map[string]any{
				"add_menu_list": &menuModel,
			}))
		return
	}

	// 遍历图片id，将图片id和菜单id添加到关联表
	for _, imageList := range menuAddParams.ImageList {
		menuImgModel = append(menuImgModel, model.MenuImagesModel{
			MenuModelID:  menuModel.ID,
			ImageModelID: imageList.ID,
		})
	}
	// 给第三张表入库
	err = M.Orm.Create(&menuImgModel).Error
	if err != nil {
		global.Mlog.Error("菜单图片关联失败:%s", zap.Error(err))
		response.Err400(
			ctx,
			response.WithMsg("菜单图片关联失败"))
		return
	}
	// 返回成功的响应
	response.OK201(
		ctx,
		response.WithMsg("菜单创建成功"),
		response.WithData(map[string]any{
			"add_menu_list": &menuModel,
		}))
}
