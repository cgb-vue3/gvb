package core

import "akita/pkg/redis"

func InitRedis() {
	redis.ConnectRedis()
}
