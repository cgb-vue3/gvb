package side_menu_dao

import (
	"akita/app/http/controllers/common"
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/paging"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (M SideMenuDao) GetPag(ctx *gin.Context, sidePagParams common.PagingParams, role int) {
	var (
		//查找到所有的主菜单切片
		sideMenu []models.SideMenuModel
		// 查找到所有子菜单切片
		childMenu []models.ChildSideMenuModel
		// 用于返回的主菜单切片
		rspSideMenuModelList = make([]models.SideMenuModel, 0)
		//用于返回的子菜单切片
		respChildMenuList = make([]models.ChildSideMenuModel, 0)
	)

	// 超级管理员 --> 给所有权限
	// 管理员 -->  给2级权限的菜单
	// 普通用户 --> 给3级权限的菜单
	// 游客   ---> 不给权限
	// 禁止用户 --> 禁止登录

	// 将与用户等级相对应的主菜单和子菜单筛选出来
	switch role {
	case 1:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 1, 3, sidePagParams)
	case 2:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 2, 3, sidePagParams)
	case 3:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 3, 3, sidePagParams)
	}
	//遍历菜单和子菜单，对比菜单和子菜单的id是否相等，如果相等就将子菜单添加到主菜单中
	for _, rootModel := range sideMenu {
		for _, childModel := range childMenu {
			if rootModel.ID == childModel.SideMenuModelID {
				// 将与主菜单关联的子菜单添加到childMenu切片中
				rootModel.ChildMenu = append(rootModel.ChildMenu, childModel)
				continue
			}
			// 将没有子菜单的二级菜单数组置为空切片
			rootModel.ChildMenu = respChildMenuList
		}
		// 将主菜单切片赋值给sideMenuModelList，用于返回
		rspSideMenuModelList = append(rspSideMenuModelList, rootModel)
	}
	response.OK200(ctx, response.WithMsg("cg"), response.WithData(map[string]any{
		"data": rspSideMenuModelList,
	}))
}

// 筛选符合用户权限级别的菜单
func screenOutMenu(ctx *gin.Context, Orm *gorm.DB, start, end int, sidePagParams common.PagingParams) ([]models.SideMenuModel, []models.ChildSideMenuModel) {
	// 查找主菜单
	var (
		// 查找失败返回空切片
		sideMenu = make([]models.SideMenuModel, 0)
		// 查找失败返回空切片
		childMenu = make([]models.ChildSideMenuModel, 0)
		// 查找到符合权限的主菜单
		sideMenuModel []models.SideMenuModel
		// 查找到符合权限的子菜单
		childMenuModel []models.ChildSideMenuModel
	)

	// 返回分页、排序后的菜单列表
	sideMenuPagList, err := paging.Pag(&sideMenuModel, Orm, paging.Option{Params: sidePagParams}, "find")
	if err != nil {
		response.Err400(
			ctx,
			response.WithMsg("查询菜单列表失败"),
			response.WithErr(err),
		)
		return sideMenu, childMenu
	}
	sideMenuModel = *sideMenuPagList

	if err := Orm.Where("level BETWEEN ? AND ?", start, end).Find(&sideMenuModel).Error; err != nil {
		global.Mlog.Error("菜单获取失败")
		response.Err400(ctx, response.WithMsg("菜单获取失败"))
		return sideMenu, childMenu
	}

	// 返回分页、排序后的菜单列表
	childMenuPagList, err := paging.Pag(&childMenuModel, Orm, paging.Option{Params: sidePagParams}, "find")
	if err != nil {
		response.Err400(
			ctx,
			response.WithMsg("查询菜单列表失败"),
			response.WithErr(err),
		)
		return sideMenu, childMenu
	}
	childMenuModel = *childMenuPagList

	// 查找所有子菜单
	if err := Orm.Where("level BETWEEN ? AND ?", start, end).Find(&childMenuModel).Error; err != nil {
		global.Mlog.Error("子菜单获取失败")
		response.Err400(ctx, response.WithMsg("菜单获取失败"))
		return sideMenu, childMenu
	}
	return sideMenuModel, childMenuModel
}
