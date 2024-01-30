package public_Routers

import (
	"akita/app/http/controllers/api/v1/public"
	"github.com/gin-gonic/gin"
)

// localhost:8080/api/v1/images
func (PublicRouters) ImageRouters(group *gin.RouterGroup, apis *public.PublicGroupApis) {
	group.POST("images", apis.Upload.Add)
	group.GET("images_pag", apis.Upload.ResponsePagingList)
	group.DELETE("images", apis.Upload.Delete)
	//group.GET("images_list", apis.Upload.ResponseImageList)
}
