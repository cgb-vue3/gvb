package global

import (
	"akita/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量
var (
	// MConfig 全局配置信息
	MConfig *conf.Config
	// MDB 全局数据库
	MDB *gorm.DB
	// Mlog 全局日志
	Mlog *zap.Logger
	//MRedis *core.RedisClient
)
