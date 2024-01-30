package side_menu_dao

import (
	"akita/app/http/controllers/api/v1/auth/side_menu_api/side_menu_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M SideMenuDao) Add(ctx *gin.Context, sideMenuParams side_menu_params.SideMenuParams) {
	// 查找数据库中是否存在该菜单
	if err := M.Orm.Where("title = ?", sideMenuParams.Title).Take(&models.SideMenuModel{}).Error; err == nil {
		global.Mlog.Error("菜单存在", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeMenuExists))
		return
	}
	//  菜单不存在就创建
	var sideMenuModel = models.SideMenuModel{
		Title: sideMenuParams.Title,
		Icon:  sideMenuParams.Icon,
		Name:  sideMenuParams.Name,
		Path:  sideMenuParams.Path,
		Level: sideMenuParams.Level,
		Sort:  sideMenuParams.Sort,
	}

	if err := M.Orm.Create(&sideMenuModel).Error; err != nil {
		global.Mlog.Error("菜单添加失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeMenuCreationFailed))
		return
	}
	response.OK201(ctx, response.CodeWithMsg(response.CodeMenuCreatedSuccess))
}
