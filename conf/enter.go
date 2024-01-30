package conf

type Config struct {
	Mysql  *Mysql
	Logger *Logger
	System *System
	QQ     *QQ
	QiNiu  *QiNiu
	Upload *Upload
	JWT    *JWT
	Redis  *Redis
	Email  *Email
}
