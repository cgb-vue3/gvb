package core

import (
	"akita/conf"
	"akita/global"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

// InitConfig viper读取yaml配置文件
func InitConfig() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("路径获取失败 [ERROR]:%s", err.Error())
	}
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path + "/")

	// 使用new实例化conf.config,并赋值给全局MConfig变量
	global.MConfig = new(conf.Config)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到错误；如果需要可以忽略
			fmt.Println("配置文件未找到错误；如果需要可以忽略")
		} else {
			// 配置文件被找到，但产生了另外的错误
			fmt.Println("配置文件未找到错误；如果需要可以忽略")
		}
	}
	err = viper.Unmarshal(global.MConfig)
	if err != nil {
		fmt.Printf("导出数据有误 [ERROR]%s", err.Error())
	}
	log.Println("配置文件初始化成功")
}
