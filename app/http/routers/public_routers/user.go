package public_Routers

import (
	"akita/app/http/controllers/api/v1/public"
	"github.com/gin-gonic/gin"
)

// UserRouters 用户路由
func (PublicRouters) UserRouters(group *gin.RouterGroup, apis *public.PublicGroupApis) {
	group.POST("register", apis.User.Register)
	group.POST("login", apis.User.Login)
	group.POST("email_code", apis.User.SendEmail)
	group.POST("change_pwd", apis.User.ChangePwd)
	/*
		其它无需验证的路由
	*/
	// 获取所有文章	localhost:8080/api/v1/article/all
	group.GET("article/all", apis.User.GetAllArticle)
}
