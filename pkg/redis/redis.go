package redis

import (
	"akita/global"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sync"
	"time"
)

// RedisClient 服务
type RedisClient struct {
	Client            *redis.Client
	Context           context.Context
	address, password string
	db                int
}

// Redis 全局redis对象
var Redis *RedisClient

// once 确保全局的 Redis 对象只实例一次
var once sync.Once

// ConnectRedis 连接 redis 数据库，设置全局的 Redis 对象
func ConnectRedis() {

	once.Do(func() {
		NewClient()
	})
}

// NewClient 创建一个新的 redis 连接
func NewClient() {
	//ctx = context.Background()
	Redis = &RedisClient{
		Context:  context.Background(),
		address:  fmt.Sprintf("%v:%v", global.MConfig.Redis.Host, global.MConfig.Redis.Port),
		password: global.MConfig.Redis.Password,
		db:       global.MConfig.Redis.DB,
	}

	// 使用 redis 库里的 NewClient 初始化连接
	Redis.Client = redis.NewClient(&redis.Options{
		Addr:     Redis.address,
		Password: Redis.password,
		DB:       Redis.db,
	})
	// 测试一下连接
	ping()
}

// Ping 用以测试 redis 连接是否正常
func ping() {

	_, err := Redis.Client.Ping(Redis.Context).Result()
	if err != nil {
		global.Mlog.Error("redis连接错误", zap.Error(err))
		return
	}
	global.Mlog.Info("redis连接成功")
}

// Set 存储 key 对应的 value，且设置 expiration 过期时间
func Set(key string, value interface{}, expiration time.Duration) bool {
	if err := Redis.Client.Set(Redis.Context, key, value, expiration).Err(); err != nil {
		global.Mlog.Error("Redis Set", zap.Error(err))
		return false
	}
	return true
}

// Get 获取 key 对应的 value
func Get(key string) string {
	result, err := Redis.Client.Get(Redis.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			global.Mlog.Error("Redis Get", zap.Error(err))
		}
		return ""
	}
	return result
}

// DelStr 清除string
func DelStr(key string) error {
	_, err := Redis.Client.Del(Redis.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			global.Mlog.Error("Redis Del", zap.Error(err))
		}
		return err
	}
	return nil
}

//// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
//func (rds RedisClient) Has(key string) bool {
//	_, err := rds.Client.Get(rds.Context, key).Result()
//	if err != nil {
//		if err != redis.Nil {
//			global.Mlog.Error("Redis Has", zap.Error(err))
//
//		}
//		return false
//	}
//	return true
//}

//func (rds RedisClient) SetH(key string, value ...interface{}) bool {
//	_, err := Redis.Client.HSet(rds.Context, key, value...).Result()
//	if err != nil {
//		if err != redis.Nil {
//			global.Mlog.Error("Redis Get", zap.Error(err))
//		}
//		return false
//	}
//	Redis.Client.Sort()
//	return true
//}

//
//func (rds RedisClient) GetH(key ...string) bool {
//
//	_, err := Redis.Client.HGet(rds.Context, key).Result()
//	if err != nil {
//		if err != redis.Nil {
//			global.Mlog.Error("Redis Get", zap.Error(err))
//		}
//		return false
//	}
//	return true
//}
