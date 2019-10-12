package daemon

import (
	"GoModDemo/util"
	"github.com/garyburd/redigo/redis"
	"time"
)

//Run 运行后台线程
func Run() {
	logger := util.InitZapLog()
	logger.Debug("启动后台进程")
	go registerTask()
}

///registerTask 注册基础数据
func registerTask() {
	logger := util.InitZapLog()

	time.Sleep(2 * time.Second)
	cache := util.NewCache()
	res, err := cache.Value("deviceId")
	if err != nil {
		logger.Error("查找缓存设备号异常,无法启动后台进程")
		return
	}
	deviceId := "deviceId:" + res.Data().(string)
	if deviceId == "" {
		logger.Error("设备号为空,无法启动后台进程")
		return
	}
	go pushTask(deviceId)
	go taskUserInfoByAccountName()
}

//PushTask 推送基础数据
func pushTask(deviceId string) {
	logger := util.InitZapLog()
	for {
		
		isOk, err := GetPushQx(deviceId)
		if err != nil {

			logger.Error("抢推送基础资料权限失败")
			return
		}
		isOk=true
		if isOk {
			logger.Debug("本机" + deviceId + "推送基础资料权限")
		}
		time.Sleep(1800* time.Second)

	}
}

//GetPushQx 查看是否有权限推送基础数据
func GetPushQx(deviceId string) (bool, error) {
	rc := util.RedisClient.Get()
	defer rc.Close()
	
	v, err := redis.Int64(rc.Do("EXISTS", "pushdeviceId"))
	if err != nil {
		return false, err
	}
	if v != 1 {
		rc.Do("SET", "pushdeviceId", deviceId, "EX", "4000")
		return true, nil
	}
	v1, err := redis.String(rc.Do("GET", "pushdeviceId"))
	if err != nil {
		return false, err
	}
	if v1 != deviceId {
		return false, nil
	}
	return true, nil
}

func taskUserInfoByAccountName() {
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()
	time.Sleep(2* time.Second)
	for {
		
		
	     go TaskUserInfoByAccountName()
		
		time.Sleep(120* time.Second)

	}

	
}

