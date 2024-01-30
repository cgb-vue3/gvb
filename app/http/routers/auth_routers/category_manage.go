package auth_routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func (AuthRouters) CategoryManage(group *gin.RouterGroup, apis *auth.AuthGroupApis) {
	//localhost:8080/api/auth/v1/category/add
	group.POST("/category/add", apis.CategoryManage.Add)
}
