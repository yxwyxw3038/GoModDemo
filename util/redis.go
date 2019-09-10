package util

import (
	"GoModDemo/setting"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	// 定义常量
	RedisClient    *redis.Pool
	REDIS_HOST     string
	REDIS_DB       string
	REDIS_PASSWORD string
)

func init() {
	setting.Setup()
	// 从配置文件获取redis的ip以及db
	REDIS_HOST = setting.RedisSetting.Url
	REDIS_DB = setting.RedisSetting.DbName
	REDIS_PASSWORD = setting.RedisSetting.Password
	// MaxIdle :=setting.RedisSetting.MaxIdle
	// MaxActive :=setting.RedisSetting.MaxActive

	// 建立连接池
	RedisClient = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     1,
		MaxActive:   10,
		IdleTimeout: 180 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", REDIS_HOST)
			if err != nil {
				return nil, err
			}
			_, err = c.Do("AUTH", REDIS_PASSWORD)
			if err != nil {
				c.Close()
				return nil, err
			}
			// 选择db
			c.Do(REDIS_DB)
			return c, nil
		},
	}
}
