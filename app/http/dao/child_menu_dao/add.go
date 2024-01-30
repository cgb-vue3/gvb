package child_menu_dao

import (
	"akita/app/http/controllers/api/v1/auth/child_menu_api/child_menu_params"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (M ChildMenuDao) Add(ctx *gin.Context, childMenuParams child_menu_params.ChildMenuParams) {

	// 查找数据库中是否存在该菜单
	if err := M.Orm.Where("title = ?", childMenuParams.Title).Take(&models.ChildSideMenuModel{}).Error; err == nil {
		global.Mlog.Error("菜单存在", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeSonMenuExists))
		return
	}

	var (
		childMenuModel = models.ChildSideMenuModel{
			Title:           childMenuParams.Title,
			Icon:            childMenuParams.Icon,
			Name:            childMenuParams.Name,
			Path:            childMenuParams.Path,
			Sort:            childMenuParams.Sort,
			Level:           childMenuParams.Level,
			SideMenuModelID: childMenuParams.ParentID,
		}
		model models.SideMenuModel
	)

	// 创建子菜单
	if err := M.Orm.Create(&childMenuModel).Error; err != nil {
		global.Mlog.Error("子菜单添加失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeSonMenuCreationFailed))
		return
	}

	// 根据传入的主菜单id，查找到主菜单
	M.Orm.Take(&model, "id = ?", childMenuParams.ParentID)
	// 更新主菜单的ChildMenu字段
	if err := M.Orm.Model(&model).Select("ChildMenu").Updates(&childMenuModel).Error; err != nil {
		global.Mlog.Error("子菜单添加失败", zap.Error(err))
		response.Err400(ctx, response.CodeWithMsg(response.CodeSonMenuCreationFailed))
		return
	}
	response.OK201(ctx, response.CodeWithMsg(response.CodeSonMenuCreatedSuccess))
}
