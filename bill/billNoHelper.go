package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	"strconv"
	// "strings"
)

func GetAllBillNoInfo(BillNoStr string, PageSize, CurrentPage int) (*[]model.BillNo, int, error) {
	list := make([]model.BillNo, 0)
	whereSql, err := util.GetWhereSqlOrderLimt("BillNo", BillNoStr, "UpdateTime", consts.DESC, PageSize, CurrentPage)
	if err != nil {
		return nil, 0, err
	}
	whereSqlCount, err := util.GetWhereSqlCount("BillNo", BillNoStr)
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
		var temp model.BillNo
		temp.ID = util.ToString(data[i]["ID"])
		temp.CodeName = util.ToString(data[i]["CodeName"])
		temp.CurrentId = util.ToInt(data[i]["CurrentId"])
		temp.Code = util.ToString(data[i]["Code"])
		temp.CurrentBillNo = util.ToString(data[i]["CurrentBillNo"])
		temp.MaskInfo = util.ToString(data[i]["MaskInfo"])
		temp.EndLength = util.ToInt(data[i]["EndLength"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		currentTime, _ := util.AnyToTimeStr(data[i]["CurrentTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.CurrentTime = currentTime
		list = append(list, temp)
	}
	return &list, num, nil
}

func GetBillNoByID(ID string) (*model.BillNo, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.BillNo
	err = db.Table(&data).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("未找到对应单据")
	}
	if len(data) > 1 {
		return nil, errors.New("找到多个对应单据")
	}
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime, _ = util.ParseAnyToStr(data[0].UpdateTime)
	data[0].CurrentTime, _ = util.ParseAnyToStr(data[0].CurrentTime)
	return &data[0], nil
}
func DeleteBillNo(idList []string) error {
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
		_, err = db.Table("BillNo").Where("ID", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}

	}
	db.Commit()
	return err
}

func AddBillNo(data model.BillNo) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.CreateTime = timeStr
	data.UpdateTime = timeStr
	data.CurrentTime = timeStr
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	count, err := db.Table("BillNo").Where("Code", "=", data.Code).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Code不能重复")
	}
	_, err = db.ExtraCols(consts.GetBillNoInfo()...).Insert(&data)
	if err != nil {
		return err
	}
	return err
}

func UpdateBillNo(data model.BillNo) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("修改数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.UpdateTime = timeStr
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.ExtraCols(consts.GetBillNoInfo()...).Where("ID", data.ID).Update(&data)
	if err != nil {
		return err
	}
	return err
}
func GetBillNo(Code string) (string, error) {
	var err error
	str := ""
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("修改数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return "", err
	}
	var data []model.BillNo
	err = db.Table(&data).Where("Code", "=", Code).Select()
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", errors.New("未找到对应单据")
	}
	if len(data) > 1 {
		return "", errors.New("找到多个对应单据")
	}
	newTime, timeStr := util.GetNowAndStr()
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime = timeStr

	oldCurrentId := data[0].CurrentId
	ymd, err := util.GetMaskDataStr(data[0].MaskInfo, newTime)
	if err != nil {
		return "", err
	}
	oldTime, err := util.ParseAny(data[0].CurrentTime)
	if err != nil {
		return "", err
	}
	ymdOld, err := util.GetMaskDataStr(data[0].MaskInfo, oldTime)
	if err != nil {
		return "", err
	}
	if ymdOld != ymd {
		data[0].CurrentId = 1
	} else {
		data[0].CurrentId++
	}
	ends := fmt.Sprintf("%0"+strconv.Itoa(data[0].EndLength)+"d", data[0].CurrentId)
	str = data[0].Code + ymd + ends
	data[0].CurrentTime = timeStr
	data[0].CurrentBillNo = str
	count, err := db.ExtraCols(consts.GetBillNoInfo()...).Where("ID", data[0].ID).Where("CurrentId", oldCurrentId).Update(&data)
	if err != nil {
		return "", err
	}
	if count <= 0 {
		return "", errors.New("生成单据号错误")
	}
	return str, nil

}
