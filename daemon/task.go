package daemon

import (
	"GoModDemo/bill"
	"GoModDemo/model"
	"GoModDemo/setting"
	"GoModDemo/util"
	"encoding/json"
	"fmt"
	"strconv"

	"go.uber.org/zap"
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
func TaskWs(logger *zap.Logger) {
	defer func() {
		if p := recover(); p != nil {
			logger.Error("异常")
		}
	}()
	logger.Sync()

	db, err := util.OpenDB()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	strSql := fmt.Sprintf("select n.ID, n2.UserId,n.TypeId,n.NoticeTime,n.No,n.Title,n.Content from Notice as n,NoticeUser as n2 where n.ID=n2.NoticeId and n.Status=5 and IFNULL( n2.SendFlag,0)=0 and  DATE_FORMAT(n.SendBeginTime,'%Y-%m-%d %H:%i')<=DATE_FORMAT(sysdate(),'%Y-%m-%d %H:%i' ) and  DATE_FORMAT(n.SendEndTime,'%Y-%m-%d %H:%i')>=DATE_FORMAT(sysdate(),'%Y-%m-%d %H:%i' )")
	data, err := db.Query(strSql)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	timeStr := util.GetNowStr()
	for k, v := range util.WSManager.Clients {
		if v {
			for i := 0; i < len(data); i++ {
				UserId := util.ToString(data[i]["UserId"])
				if (*k).UserId == UserId {
					if err != nil {
						logger.Error(err.Error())
						break
					}
					var temp model.MsgInfo
					temp.ID = util.ToString(data[i]["ID"])
					temp.No = util.ToString(data[i]["No"])
					temp.TypeId = util.ToInt(data[i]["TypeId"])
					temp.Title = util.ToString(data[i]["Title"])
					temp.Content = util.ToString(data[i]["Content"])
					noticeTime, _ := util.AnyToTimeStr(data[i]["NoticeTime"])
					temp.BillTime = noticeTime
					temp.MsgType = "Notice"
					temp.OpenFlag = 0
					temp.MsgTime = timeStr
					b, err := json.Marshal(temp)
					if err != nil {
						logger.Error(err.Error())
						break
					}

					msg := util.WSMessage{ID: (*k).ID, Type: "Msg", Data: string(b)}
					b, err = json.Marshal(msg)
					if err != nil {
						logger.Error(err.Error())
						break
					}

					util.WSManager.Send(b, k)
					err = bill.UpdateNoticeUserStatus(temp.ID, UserId)
					if err != nil {
						logger.Error(err.Error())
						break
					}
				}

			}
		}

	}
	logger.Debug("结束写入参数缓存")
}
