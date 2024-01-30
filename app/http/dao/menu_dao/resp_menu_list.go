package menu_dao

import (
	"akita/app/http/controllers/api/v1/public/menu_api/menu_resp"
	model "akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ResponseMenuList 返回菜单列表
func (M MenuDao) ResponseMenuList(ctx *gin.Context, respMenuList []menu_resp.RespMenuList) {
	var (
		menuModel      []model.MenuModel
		menuIdList     []uint
		menuImageModel []model.MenuImagesModel
	)
	// 先拿到菜单id
	err := M.Orm.Order("sort").Find(&menuModel).Select("id").Scan(&menuIdList).Error
	if err != nil {
		global.Mlog.Error("查询menuID失败:", zap.Error(err))
		return
	}

	// 拿菜单id去查第三张关联表，将查到的记录添加到menuImageModel切片中
	err = M.Orm.Order("image_model_id desc").Preload("ImageModel").Find(&menuImageModel, "menu_model_id in ?", menuIdList).Error
	if err != nil {
		global.Mlog.Error("连表查询失败:", zap.Error(err))
		return
	}

	// 遍历menuModelList,判断menu的id与关联表中的id是否相等。不相等就跳出循环，相等就把关联表中的图片id添加到Images。
	//再将menu和Images添加到responseMenusModel。
	for _, menu := range menuModel {
		imageList := []menu_resp.RespImageList{}
		for _, menuImage := range menuImageModel {
			if menu.ID != menuImage.MenuModelID {
				continue
			}
			imageList = append(imageList, menu_resp.RespImageList{
				ID:   menuImage.ImageModelID,
				Name: menuImage.ImageModel.Name,
				Path: menuImage.ImageModel.Path,
			})
		}
		respMenuList = append(respMenuList, menu_resp.RespMenuList{
			MenuModel: menu,
			Images:    imageList,
		})
	}
	response.OK200(
		ctx,
		response.WithData(map[string]any{
			"resp_menu_list": respMenuList,
		}))
}
