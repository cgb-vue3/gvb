package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) TagManage(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/tag/add
	group.POST("/tag/add", apis.TagManage.Add)
}
