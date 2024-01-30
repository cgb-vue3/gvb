package core

import (
	"akita/app/http/routers"
	"akita/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// InitServer 初始化gin
func InitServer() {
	// 设置为ReleaseMode模式
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	// 配置中间件
	//RegisterGlobalMiddleWare(r)
	// 注册路由
	routers.RegisterRouterApi(r)
	// 启动服务
	startUpServer(r)
}

//func RegisterGlobalMiddleWare(router *gin.Engine) {
//	// 使用自定义的logger、Recovery中间件
//	router.Use(
//		middleware.ZapLogger(global.Mlog),
//		middleware.ZapRecovery(global.Mlog, true),
//	)
//}

func startUpServer(router *gin.Engine) {
	// 端口
	addr := global.MConfig.System.Addr()
	if addr != "" {
		addr = ":8080"
	}
	global.Mlog.Info(fmt.Sprintf("Listening and serving HTTP on %s", addr))

	err := router.Run(addr)
	if err != nil {
		global.Mlog.Error("服务器启动失败:", zap.Error(err))
		return
	}
}
