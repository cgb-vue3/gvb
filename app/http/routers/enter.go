package routers

import (
	"akita/app/http/controllers/api/v1/auth"
	"akita/app/http/controllers/api/v1/public"
	"akita/app/http/middleware"
	"akita/app/http/routers/auth_routers"
	"akita/app/http/routers/public_routers"
	"akita/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

// Routes 用于接收路由的结构体
type Routes struct {
	PublicRouters public_Routers.PublicRouters
	AuthRouters   auth_routers.AuthRouters
}

// GroupRoutes 实例化Routes
var GroupRoutes = new(Routes)

func RegisterRouterApi(r *gin.Engine) {
	r.Use(
		middleware.ZapLogger(global.Mlog),
		middleware.ZapRecovery(global.Mlog, true),
	)

	// 添加swagger访问路由
	// 构建您的应用程序，然后转到 http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.Use(middleware.CorsAuthMiddleware())
	// 实例化公共分组的api
	// localhost:8080/api/v1
	var publicApi = public.NewPublicApis()
	groupPublic := r.Group("/api")
	publicGroup := groupPublic.Group("v1")
	{
		//上传图片
		GroupRoutes.PublicRouters.ImageRouters(publicGroup, publicApi)
		// 前台菜单
		GroupRoutes.PublicRouters.MenuRouters(publicGroup, publicApi)
		// 用户
		GroupRoutes.PublicRouters.UserRouters(publicGroup, publicApi)
	}

	// 需要鉴权的路由
	// localhost:8080/api_auth/v1
	var authApi = auth.NewAuthApis()
	groupAuth := r.Group("/api/auth", middleware.JWTAuthMiddleware(), middleware.CorsAuthMiddleware())
	authGroup := groupAuth.Group("v1")
	{
		// 用户管理
		GroupRoutes.AuthRouters.UserManage(authGroup, authApi)
		// 后台菜单
		GroupRoutes.AuthRouters.SideMenu(authGroup, authApi)
		// 后台子管理
		GroupRoutes.AuthRouters.ChildMenu(authGroup, authApi)
		// 文章管理
		GroupRoutes.AuthRouters.ArticleManage(authGroup, authApi)
		// 分类管理
		GroupRoutes.AuthRouters.CategoryManage(authGroup, authApi)
		// 标签管理
		GroupRoutes.AuthRouters.TagManage(authGroup, authApi)
		// 评论管理
		GroupRoutes.AuthRouters.CommentManage(authGroup, authApi)
	}
}
