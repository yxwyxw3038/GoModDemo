package daemon

import (
	"GoModDemo/setting"
	"GoModDemo/util"
	"github.com/garyburd/redigo/redis"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type DataSyncMutex struct {
	mutex  sync.RWMutex
	isSync bool
}

var dataSync DataSyncMutex

//Run 运行后台线程
func Run() {
	logger := util.InitZapLog()
	logger.Debug("启动后台进程")
	// go registerTask()
	// go jumpTask()
	dataSync.mutex.Lock()
	dataSync.isSync = false
	dataSync.mutex.Unlock()
	var wait util.WaitGroupWrapper
	wait.Wrap(registerTask)
	wait.Wrap(jumpTask)
	wait.Wait()
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
	go taskRun()
}
func jumpTask() {
	logger := util.InitZapLog()
	times := setting.AppSetting.JumpTime
	for {
		logger.Sync()
		count := runtime.NumGoroutine()
		logger.Debug("心跳服务，总运行Goroutine数量" + strconv.Itoa(count))
		time.Sleep(time.Second * time.Duration(times))
	}
}

//PushTask 推送基础数据
func pushTask(deviceId string) {
	defer dataSync.mutex.Unlock()
	times := time.Duration(setting.AppSetting.TaskTime)
	logger := util.InitZapLog()
	run := func() {

		isOk, err := GetPushQx(deviceId)
		if err != nil {

			logger.Error("抢推送基础资料权限失败")
			return
		}
		isOk = true
		if isOk {
			logger.Debug("本机" + deviceId + "推送基础资料权限")
		}
		dataSync.mutex.Lock()
		dataSync.isSync = isOk
		dataSync.mutex.Unlock()
		// time.Sleep(times * time.Second)

	}
	run()
	ticker := time.NewTicker(times * time.Second)
	for {
		select {
		case <-ticker.C:
			run()
		}
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
func taskRun() {
	times := time.Duration(setting.AppSetting.DelayTime)
	time.Sleep(times * time.Second)
	var wait util.WaitGroupWrapper
	wait.Wrap(taskUserInfoByAccountName)
	wait.Wrap(taskParameter)
	wait.Wrap(taskWs)
	wait.Wait()
}
func taskUserInfoByAccountName() {
	defer dataSync.mutex.RUnlock()
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()

	times := time.Duration(setting.AppSetting.TaskTime1)
	// for {
	// 	dataSync.mutex.RLock()
	// 	isSync := dataSync.isSync
	// 	dataSync.mutex.RUnlock()
	// 	if isSync {
	// 		TaskUserInfoByAccountName(logger)
	// 	}

	// 	time.Sleep(times * time.Second)

	// }
	run := func() {
		dataSync.mutex.RLock()
		isSync := dataSync.isSync
		dataSync.mutex.RUnlock()
		if isSync {
			TaskUserInfoByAccountName(logger)
		}
	}
	run()
	ticker := time.NewTicker(times * time.Second)
	for {
		select {
		case <-ticker.C:
			run()
		}
	}

}

func taskParameter() {
	defer dataSync.mutex.RUnlock()
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()

	times := time.Duration(setting.AppSetting.TaskTime1)
	// for {
	// 	dataSync.mutex.RLock()
	// 	isSync := dataSync.isSync
	// 	dataSync.mutex.RUnlock()
	// 	if isSync {
	// 		TaskUserInfoByAccountName(logger)
	// 	}

	// 	time.Sleep(times * time.Second)

	// }
	run := func() {
		dataSync.mutex.RLock()
		isSync := dataSync.isSync
		dataSync.mutex.RUnlock()
		if isSync {
			TaskParameter(logger)
		}
	}
	run()
	ticker := time.NewTicker(times * time.Second)
	for {
		select {
		case <-ticker.C:
			run()
		}
	}

}

func taskWs() {
	defer dataSync.mutex.RUnlock()
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()

	times := time.Duration(setting.AppSetting.TaskTime2)

	run := func() {
		dataSync.mutex.RLock()
		isSync := dataSync.isSync
		dataSync.mutex.RUnlock()
		if isSync {
			TaskWs(logger)
		}
	}
	run()
	ticker := time.NewTicker(times * time.Second)
	for {
		select {
		case <-ticker.C:
			run()
		}
	}

}
