package core

import (
	"akita/app/http/models"
	"akita/global"
	lumberjack2 "akita/pkg/lumberjack"
	"errors"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

// InitDB 连接数据库
func InitDB() {
	switch global.MConfig.Mysql.DataBase {
	case "mysql":
		//获取配置信息
		// 连接数据库的dsn
		dsn := global.MConfig.Mysql.Dsn()
		// 最大空闲连接数
		MaxIdleConns := global.MConfig.Mysql.MaxIdleConns
		// 最大连接数
		MaxOpenConns := global.MConfig.Mysql.MaxOpenConns

		logLevel := setLogLevel()

		DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logLevel,
		})

		if err != nil {
			global.Mlog.Fatal("数据库连接失败:", zap.Error(err))
		}
		global.MDB = DB
		global.Mlog.Info("数据库连接成功")

		// 连接池配置
		sqlDB, _ := global.MDB.DB()
		// 最大空闲连接数
		sqlDB.SetMaxIdleConns(MaxIdleConns)
		// 最大打开连接数
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		// 连接超时时间
		sqlDB.SetConnMaxLifetime(time.Hour)

		// 创建多表连接
		err = global.MDB.SetupJoinTable(&models.MenuModel{}, "ImageModel", &models.MenuImagesModel{})
		if err != nil {
			global.Mlog.Error("多表连接失败", zap.Error(err))
			return
		}
		err = global.MDB.SetupJoinTable(&models.ArticleModel{}, "TagModel", &models.ArticleTagModel{})
		if err != nil {
			global.Mlog.Error("多表连接失败", zap.Error(err))
			return
		}
		//AutoMigrate 用于自动迁移表结构
		err = global.MDB.Set("gorm:table_options", "ENGINE=InnoDB").
			AutoMigrate(
				&models.UserModel{},
				&models.ImageModel{},
				&models.MenuModel{},
				&models.MenuImagesModel{},
				&models.SideMenuModel{},
				&models.ChildSideMenuModel{},
				&models.ArticleTagModel{},
				&models.CategoryModel{},
				&models.ArticleModel{},
				&models.TagModel{},
				&models.CommentModel{},
			)
		if err != nil {
			global.Mlog.Fatal("数据表迁移失败:", zap.Error(err))
		}
		global.Mlog.Info("数据表迁移成功")

	default:
		panic(errors.New("database connection not supported"))
	}
}

// 设置日志级别
func setLogLevel() logger.Interface {
	var logLevel logger.LogLevel
	var newLog *log.Logger
	enable := global.MConfig.Mysql.LogEnable
	level := global.MConfig.Mysql.LogLevel
	logType := global.MConfig.Mysql.LogType
	logPath := global.MConfig.Mysql.LogPath

	filePath := logPath + time.Now().Format("2006-01-02") + ".txt"

	switch enable {
	case level == "Silent":
		logLevel = logger.Silent
	case level == "Error":
		logLevel = logger.Error
	case level == "warn":
		logLevel = logger.Warn
	case level == "info":
		logLevel = logger.Info
	}

	logJack := lumberjack2.LJackLog(filePath)
	if logType == "files" {
		writer := io.MultiWriter(logJack, os.Stdout)
		newLog = log.New(writer, "\r\n", log.LstdFlags)
	} else {
		newLog = log.New(os.Stdout, "\r\n", log.LstdFlags)
	}

	newLogger := logger.New(
		newLog,
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	return newLogger
}
