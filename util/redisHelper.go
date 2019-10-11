package util
import (
	"github.com/garyburd/redigo/redis"
	"errors"
)
func RedisExists(key string) (bool,  error) {
	rc := RedisClient.Get()
	defer rc.Close()
	
	v, err := redis.Int64(rc.Do("EXISTS", key))
	if err != nil {
		return  false, err
	}
	if v != 1 {
		return  false,nil
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

func SetRedisAny(key string, any interface{}) ( error) {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("设置Redis Key："+key+"异常")
		}
	}()
	rc.Do("SET", key, any)

	return err
}
func SetRedisAnyEx(key string, any interface{},exTime string) ( error) {
	rc := RedisClient.Get()
	var err error
	defer rc.Close()
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("设置Redis Key："+key+"异常")
		}
	}()
	rc.Do("SET", key, any, "EX", exTime)

	return err
}