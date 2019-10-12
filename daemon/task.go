package daemon
import (
	"GoModDemo/util"
	"GoModDemo/model"
	"encoding/json"
)
func TaskUserInfoByAccountName()  {
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()
	hasKey:="TaskUserInfoByAccountName"

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
	hasmap :=make(map[string]string)
    for i:=0;i<len(user);i++ {
		tempStr:=""
		b, err := json.Marshal(user[i])
		if err != nil {
		
			logger.Error("用户"+user[i].AccountName+"缓存字符串生成错误"+err.Error())
			continue;
		}
		tempStr = string(b)
		hasmap[user[i].AccountName]=tempStr
	}
	util.SetRedisHasEx(hasKey,hasmap,"150")
	logger.Debug("结束写入用户信息缓存")
	
}