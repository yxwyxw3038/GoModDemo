package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"fmt"
	// "strings"
)

func GetAllNoticeInfo(BillNoStr string, PageSize, CurrentPage int) (*[]model.NoticeViewModel, int, error) {
	list := make([]model.NoticeViewModel, 0)
	whereSql, err := util.GetWhereSqlOrderLimt("NoticeView", BillNoStr, "UpdateTime", consts.DESC, PageSize, CurrentPage)
	if err != nil {
		return nil, 0, err
	}
	whereSqlCount, err := util.GetWhereSqlCount("Notice", BillNoStr)
	if err != nil {
		return nil, 0, err
	}

	fmt.Println(whereSqlCount)
	fmt.Println(whereSql)
	db, err := util.OpenDB()
	if err != nil {
		return nil, 0, err
	}

	dataCount, err := db.Query(whereSqlCount)
	if err != nil {
		return nil, 0, err
	}
	if len(dataCount) <= 0 {
		return &list, 0, nil
	}
	num := util.ToInt(dataCount[0]["Num"])
	data, err := db.Query(whereSql)
	if err != nil {
		return nil, 0, err
	}

	if len(data) <= 0 {
		return &list, 0, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.NoticeViewModel
		temp.ID = util.ToString(data[i]["ID"])
		temp.No = util.ToString(data[i]["No"])
		temp.TypeId = util.ToInt(data[i]["TypeId"])
		temp.Title = util.ToString(data[i]["Title"])
		temp.Content = util.ToString(data[i]["Content"])
		temp.Notes = util.ToString(data[i]["Notes"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		noticeTime, _ := util.AnyToTimeStr(data[i]["NoticeTime"])
		sendBeginTime, _ := util.AnyToTimeStr(data[i]["SendBeginTime"])
		sendEndTime, _ := util.AnyToTimeStr(data[i]["SendEndTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.NoticeTime = noticeTime
		temp.SendBeginTime = sendBeginTime
		temp.SendEndTime = sendEndTime
		temp.Status = util.ToInt(data[i]["Status"])
		temp.SendCount = util.ToInt(data[i]["SendCount"])
		temp.Count = util.ToInt(data[i]["Count"])
		list = append(list, temp)
	}
	return &list, num, nil
}
