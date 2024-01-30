package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) ChildMenu(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/child_menu/add
	group.POST("/child_menu/add", apis.ChildMenu.Add)
}
