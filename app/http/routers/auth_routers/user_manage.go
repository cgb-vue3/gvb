package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) UserManage(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	// localhost:8080/api/auth/v1/user_manage/list	返回用户列表
	group.GET("user_manage/list", apis.UserManage.UserPagList)
	//localhost:8080/api/auth/v1/user_manage/info	返回用户信息
	group.GET("/user_manage/info", apis.UserManage.GetInfo)
	//localhost:8080/api/auth/v1/user_manage/addUser	添加用户
	group.POST("/user_manage/addUser", apis.UserManage.Add)
	//localhost:8080/api/auth/v1/user_manage/delUser	删除用户
	group.DELETE("/user_manage/delUser", apis.UserManage.Del)
	//localhost:8080/api/auth/v1/user_manage/editUser	编辑用户
	group.PUT("/user_manage/editUser", apis.UserManage.Put)
	//localhost:8080/api/auth/v1/user_manage/getTotal	获取总数
	group.GET("/user_manage/getTotal", apis.UserManage.GetTotal)
}
