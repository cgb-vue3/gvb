package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) SideMenu(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/side_menu/add
	group.POST("/side_menu/add", apis.SideMenu.Add)
	//localhost:8080/api/auth/v1/side_menu/getPag
	group.GET("/side_menu/getPag", apis.SideMenu.GetPag)
}
