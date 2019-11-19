package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	"strings"
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
func AddNotice(data model.NoticeBillModel) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.Main.CreateTime = timeStr
	data.Main.UpdateTime = timeStr
	data.Main.NoticeTime = timeStr

	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	billNo, err := GetBillNo("TZ")
	if err != nil {
		return err
	}
	data.Main.No = billNo
	db.Begin()
	_, err = db.Table("Notice").ExtraCols(consts.GetNoticeInfo()...).Insert(&(data.Main))
	if err != nil {
		db.Rollback()
		return err
	}
	for i := 0; i < len(data.Item); i++ {
		item := new(model.NoticeUser)
		item.ID = data.Item[i].ID
		item.NoticeId = data.Item[i].NoticeId
		item.UserId = data.Item[i].UserId
		item.Notes = data.Item[i].Notes
		item.CreateBy = data.Item[i].CreateTime
		item.CreateTime = timeStr
		item.UpdateBy = data.Item[i].UpdateBy
		item.UpdateTime = timeStr
		item.UpdateBy = data.Item[i].UpdateBy
		item.SendTime = timeStr
		item.SendFlag = 0
		_, err = db.Table("NoticeUser").ExtraCols(consts.GetNoticeUserInfo()...).Insert(item)
		if err != nil {
			db.Rollback()
			return err
		}
	}

	db.Commit()
	return err

}

func UpdateNotice(data model.NoticeBillModel) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.Main.UpdateTime = timeStr
	db, err := util.OpenDB()
	delSql := ""
	if err != nil {
		return err
	}
	strSql := fmt.Sprintf("select ID from NoticeUser where NoticeId='%s' ", data.Main.ID)
	itemData, err := db.Query(strSql)
	if err != nil {
		return err
	}
	delList := make([]string, 0)
	for i := 0; i < len(itemData); i++ {
		findInfo := false
		for j := 0; j < len(data.Item); j++ {
			if util.ToString(itemData[i]["ID"]) == data.Item[j].ID {
				findInfo = true
				break
			}
		}
		if !findInfo {
			delList = append(delList, "'"+util.ToString(itemData[i]["ID"])+"'")
		}
	}
	if len(delList) > 0 {
		delwhereSql := strings.Join(delList, ",")
		delSql = fmt.Sprintf("Delete from NoticeUser where ID in(%s)", delwhereSql)
	}
	db.Begin()
	if len(delList) > 0 {
		_, err = db.Execute(delSql)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	_, err = db.Table("Notice").ExtraCols(consts.GetNoticeInfo()...).Where("ID", "=", data.Main.ID).Update(&(data.Main))
	if err != nil {
		db.Rollback()
		return err
	}
	for i := 0; i < len(data.Item); i++ {
		var item model.NoticeUser
		item.ID = data.Item[i].ID
		item.NoticeId = data.Item[i].NoticeId
		item.UserId = data.Item[i].UserId
		item.Notes = data.Item[i].Notes
		item.CreateBy = data.Item[i].CreateBy
		if data.Item[i].CreateTime == "" {
			item.CreateTime = timeStr
		} else {
			item.CreateTime = data.Item[i].CreateTime
		}
		if data.Item[i].UpdateTime == "" {
			item.UpdateTime = timeStr
		} else {
			item.UpdateTime = data.Item[i].UpdateTime
		}
		item.UpdateBy = data.Item[i].UpdateBy
		if data.Item[i].SendTime == "" {
			item.SendTime = timeStr
		} else {
			item.SendTime = data.Item[i].SendTime
		}

		item.SendFlag = data.Item[i].SendFlag
		count, err := db.Table("NoticeUser").ExtraCols(consts.GetNoticeUserInfo()...).Where("ID", "=", data.Item[i].ID).Update(&item)
		if err != nil {
			db.Rollback()
			return err
		}
		if count <= 0 {
			_, err = db.Table("NoticeUser").ExtraCols(consts.GetNoticeUserInfo()...).Insert(&item)
			if err != nil {
				db.Rollback()
				return err
			}
		}

	}

	db.Commit()
	return err

}
func DeleteNotice(idList []string) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("删除数据异常")
		}
	}()
	if len(idList) <= 0 {
		return nil
	}

	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	db.Begin()
	for _, v := range idList {
		count, err := db.Table("Notice").Where("ID", v).Where("Status", "=", 0).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		if count > 0 {
			_, err = db.Table("NoticeUser").Where("NoticeId", v).Delete()
			if err != nil {
				db.Rollback()
				return err
			}
		}
	}
	db.Commit()
	return err
}

func GetNoticeByID(ID string) (*model.Notice, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("删除数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.Notice
	err = db.Table(&data).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("未找到对应菜单")
	}
	if len(data) > 1 {
		return nil, errors.New("找到多个对应菜单")
	}
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime, _ = util.ParseAnyToStr(data[0].UpdateTime)
	data[0].NoticeTime, _ = util.ParseAnyToStr(data[0].NoticeTime)
	data[0].SendBeginTime, _ = util.ParseAnyToStr(data[0].SendBeginTime)
	data[0].SendEndTime, _ = util.ParseAnyToStr(data[0].SendEndTime)
	return &data[0], nil
}
func GetNoticeItemByID(ID string) (*[]model.NoticeUserView, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("删除数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	data := make([]model.NoticeUserView, 0)
	err = db.Table(&data).Where("NoticeId", "=", ID).Select()
	sqlLog := db.LastSql()
	fmt.Printf(sqlLog)
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return &data, err
	}
	for i := 0; i < len(data); i++ {
		data[i].CreateTime, _ = util.ParseAnyToStr(data[i].CreateTime)
		data[i].UpdateTime, _ = util.ParseAnyToStr(data[i].UpdateTime)
		data[i].SendTime, _ = util.ParseAnyToStr(data[i].SendTime)
	}

	return &data, err
}
func UpdateNoticeStatus(ID, UpdateBy string, oldStatus, newStatus int64) error {
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	timeStr := util.GetNowStr()
	count, err := db.Table("Notice").Data(map[string]interface{}{"Status": newStatus, "UpdateBy": UpdateBy, "UpdateTime": timeStr}).Where("ID", "=", ID).Where("Status", "=", oldStatus).Update()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("数据已修改")
	}
	return err
}
func UpdateNoticeUserStatus(ID, UserId string) error {
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	timeStr := util.GetNowStr()
	count, err := db.Table("NoticeUser").Data(map[string]interface{}{"SendFlag": 1, "SendTime": timeStr}).Where("NoticeId", "=", ID).Where("UserId", "=", UserId).Update()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("数据已修改")
	}
	return err
}
