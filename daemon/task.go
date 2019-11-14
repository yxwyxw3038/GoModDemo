package daemon

import (
	"GoModDemo/model"
	"GoModDemo/setting"
	"GoModDemo/util"
	"encoding/json"
	"go.uber.org/zap"
	"strconv"
)

func TaskUserInfoByAccountName(logger *zap.Logger) {
	// logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()
	logger.Sync()
	hasKey := "TaskUserInfoByAccountName"
	times := strconv.FormatInt(int64(setting.AppSetting.TaskTime1+setting.AppSetting.OffsetTime), 10)
	// isOk,err:= util.RedisExists(hasKey)
	// if err != nil {
	// 	logger.Error(err.Error())
	// 	return
	// }
	logger.Debug("开始写入用户信息缓存")
	db, err := util.OpenDB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var user []model.User
	err = db.Table(&user).Select()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if len(user) <= 0 {
		logger.Error("未找到对应用户")
		return
	}
	hasmap := make(map[string]string)
	for i := 0; i < len(user); i++ {
		tempStr := ""
		b, err := json.Marshal(user[i])
		if err != nil {

			logger.Error("用户" + user[i].AccountName + "缓存字符串生成错误" + err.Error())
			continue
		}
		tempStr = string(b)
		hasmap[user[i].AccountName] = tempStr
	}

	util.SetRedisHasEx(hasKey, hasmap, times)
	logger.Debug("结束写入用户信息缓存")

}
func TaskParameter(logger *zap.Logger) {
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()
	logger.Sync()
	Key := "TaskParameter"
	times := strconv.FormatInt(int64(setting.AppSetting.TaskTime1+setting.AppSetting.OffsetTime), 10)
	logger.Debug("开始写入参数缓存")
	db, err := util.OpenDB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var data []model.Parameter
	err = db.Table(&data).Select()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	s := string(b)
	err = util.SetRedisAnyEx(Key, s, times)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Debug("结束写入参数缓存")
}
