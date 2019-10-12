package util

import (
	"errors"
	"github.com/garyburd/redigo/redis"
)

func RedisExists(key string) (bool, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	v, err := redis.Int64(rc.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	if v != 1 {
		return false, nil
	}
	return true, nil
}

func GetRedisInt64(key string) (int64, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	v, err := redis.Int64(rc.Do("GET", key))
	if err != nil {
		return 0, err
	}
	return v, nil
}
func GetRedisString(key string) (string, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	v, err := redis.String(rc.Do("GET", key))
	if err != nil {
		return "", err
	}

	return v, nil
}
func GetRedisAny(key string) (interface{}, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	v, err := rc.Do("GET", key)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func SetRedisAny(key string, any interface{}) error {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("设置Redis Key：" + key + "异常")
		}
	}()
	rc.Do("SET", key, any)

	return err
}
func SetRedisAnyEx(key string, any interface{}, exTime string) error {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("设置Redis Key：" + key + "异常")
		}
	}()
	rc.Do("SET", key, any, "EX", exTime)

	return err
}


func SetRedisHas(key string, has map[string]string ) ( error) {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("设置Redis Key："+key+"异常")
		}
	}()
	rc.Do("del", key)
	for k, v := range has {
		rc.Do("hmset", key, k,v)
	}
	return err
}
func SetRedisHasEx(key string, has map[string]string ,exTime string) ( error) {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("设置Redis Key："+key+"异常")
		}
	}()
	rc.Do("del", key)
	for k, v := range has {
		rc.Do("hmset", key, k,v)
	}
	rc.Do("Expire", key, exTime)
	return err
}
func GetRedisHasString(haskey string,mapkey string) (string, error) {
	rc := RedisClient.Get()
	defer rc.Close()

	v, err := redis.String(rc.Do("HGET", haskey,mapkey))
	if err != nil {
		return "", err
	}
	return v, nil
}