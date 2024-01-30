package public_Routers

import (
	"akita/app/http/controllers/api/v1/public"
	"github.com/gin-gonic/gin"
)

// localhost:8080/api/v1/menu
func (M PublicRouters) MenuRouters(group *gin.RouterGroup, apis *public.PublicGroupApis) {
	group.POST("menu", apis.Menu.Add)
	group.GET("menu", apis.Menu.ResponseMenuList)
	group.PUT("menu", apis.Menu.Put)
	group.DELETE("menu", apis.Menu.Delete)
}
