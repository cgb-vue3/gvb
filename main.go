package main

import (
	"akita/app/core"
)

// @title			博客系统
// @version		1.0
// @description	接口文档
// @termsOfService	https://gitee.com/QTAkita/gvb
// @contact.name	Akita
// @contact.email	1240092443@qq.com
// @host			localhost:8080
func main() {
	//初始化配置文件
	core.InitConfig()
	// 初始化日志
	core.InitLogger()
	// 初始化数据库
	core.InitDB()
	//// redis
	//core.InitRedis()
	//初始化并启动server
	core.InitServer()
}
