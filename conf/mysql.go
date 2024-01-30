package conf

import (
	"fmt"
)

type Mysql struct {
	DataBase     string
	User         string
	Pass         string
	Host         string
	Dbname       string
	Charset      string
	LogLevel     string
	LogFormat    string
	LogPath      string
	LogType      string
	Env          string
	MaxIdleConns int
	MaxOpenConns int
	Port         int
	LogEnable    bool
}

// Dsn 配置连接数据库需要的dsn
func (MQ Mysql) Dsn() string {
	// 配置dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?%s&parseTime=True&loc=Local",
		MQ.User,
		MQ.Pass,
		MQ.Host,
		MQ.Port,
		MQ.Dbname,
		MQ.Charset,
	)

	return dsn
}
