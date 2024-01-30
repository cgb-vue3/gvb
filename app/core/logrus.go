package core

import (
	"akita/global"
	lumberjack2 "akita/pkg/lumberjack"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func InitLogger() {
	// 调用getEncoder,返回一个json对象
	writeSyncer := getLogWriter()
	// 日志写入路径
	encoder := getEncoder()
	// 判断日志级别
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(global.MConfig.Logger.Level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改根目录下settings.yaml文件中logger下level配置项")
	}
	// 创建核心
	core := zapcore.NewCore(encoder, writeSyncer, logLevel)

	// 创建logger
	global.Mlog = zap.New(
		core,
		zap.AddCaller(), // 输出路径和行号
	)
	global.Mlog.Info("日志初始化成功")
}

// 配置编码器
func getEncoder() zapcore.Encoder {
	// 创建编码器配置实例
	encoderConfig := zap.NewProductionEncoderConfig()
	// 配置编码器
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("[2006-01-02 15:04:05]")
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 日志记录介质。介质分为文件和OS.stdout
func getLogWriter() zapcore.WriteSyncer {
	// 自定义文件名
	filePath := global.MConfig.Logger.FilePath

	filename := filePath + time.Now().Format("2006-01-02") + ".log"
	// 配置日志滚动
	// 详见 ./settings.yaml中的日志配置

	log := lumberjack2.LJackLog(filename)

	//配置介质输出
	if global.MConfig.System.Env == "local" {
		// 输出到文件和终端
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(log))
	} else {
		// 仅输出到文件
		return zapcore.AddSync(log)
	}
}
