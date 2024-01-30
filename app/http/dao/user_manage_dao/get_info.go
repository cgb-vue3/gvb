package user_manage_dao

import (
	"akita/app/http/models"
	"akita/global"
	"akita/pkg/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (M UserManageDao) GetInfo(ctx *gin.Context, id uint) {
	var (
		// 用户信息
		userInfo models.UserModel
		//查找到所有的主菜单切片
		sideMenu []models.SideMenuModel
		// 查找到所有子菜单切片
		childMenu []models.ChildSideMenuModel
		// 用于返回的主菜单切片
		rspSideMenuModelList = make([]models.SideMenuModel, 0)
		//用于返回的子菜单切片
		respChildMenuList = make([]models.ChildSideMenuModel, 0)
	)

	M.Orm.Where("id = ?", id).Take(&userInfo)

	// 超级管理员 --> 给所有权限
	// 管理员 -->  给2级权限的菜单
	// 普通用户 --> 给3级权限的菜单
	// 游客   ---> 不给权限
	// 禁止用户 --> 禁止登录

	// 将与用户等级相对应的主菜单和子菜单筛选出来
	switch userInfo.Role {
	case 1:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 1, 3)
	case 2:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 2, 3)
	case 3:
		sideMenu, childMenu = screenOutMenu(ctx, M.Orm, 3, 3)
	}

	for _, root := range sideMenu {
		for _, son := range childMenu {
			if root.ID == son.SideMenuModelID {
				root.ChildMenu = append(root.ChildMenu, son)
				continue
			}
		}
		// 将没有子菜单的ChildMenu置为空切片
		if len(root.ChildMenu) == 0 {
			root.ChildMenu = respChildMenuList
		}
		rspSideMenuModelList = append(rspSideMenuModelList, root)
	}

	response.OK200(ctx,
		response.CodeWithMsg(response.CodeGetUserInfoSuccess),
		response.WithData(map[string]any{
			"userInfo": userInfo,
			"menuList": rspSideMenuModelList,
		}))
}

// 筛选符合用户权限级别的菜单
func screenOutMenu(ctx *gin.Context, Orm *gorm.DB, start, end int) ([]models.SideMenuModel, []models.ChildSideMenuModel) {
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

	if err := Orm.Where("level BETWEEN ? AND ?", start, end).Order("sort").Find(&sideMenuModel).Error; err != nil {
		global.Mlog.Error("菜单获取失败")
		response.Err400(ctx, response.CodeWithMsg(response.CodeGetUserInfoFailed))
		return sideMenu, childMenu
	}

	// 查找所有子菜单
	if err := Orm.Where("level BETWEEN ? AND ?", start, end).Order("sort").Find(&childMenuModel).Error; err != nil {
		global.Mlog.Error("子菜单获取失败")
		response.Err400(ctx, response.CodeWithMsg(response.CodeGetUserInfoFailed))
		return sideMenu, childMenu
	}
	return sideMenuModel, childMenuModel
}
