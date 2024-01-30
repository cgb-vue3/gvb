package lumberjack

import (
	"akita/global"
	"github.com/natefinch/lumberjack"
)

func LJackLog(filePath string) *lumberjack.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,                         // 日志文件路径
		MaxSize:    global.MConfig.Logger.MaxSize,    // 日志文件保留最大尺寸（M），超限后开始分割
		MaxBackups: global.MConfig.Logger.MaxBackups, // 保留日志文件最大个数
		MaxAge:     global.MConfig.Logger.MaxAge,     // 保留日志文件最大天数
	}
	return lumberJackLogger
}
