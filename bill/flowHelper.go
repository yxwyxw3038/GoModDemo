package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	"strings"

	// "strings"
	"github.com/yxwyxw3038/whysql"
)

func GetAllFlowInfo(Str string, PageSize, CurrentPage int) (*[]model.FlowView, int, error) {
	list := make([]model.FlowView, 0)
	// whereSql, err := util.GetWhereSqlOrderLimt("FlowView", Str, "UpdateTime", consts.DESC, PageSize, CurrentPage)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// whereSqlCount, err := util.GetWhereSqlCount("FlowView", Str)
	// if err != nil {
	// 	return nil, 0, err
	// }
	sqldb, err := whysql.NewWhy(Str)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSql, err := sqldb.SetTabName("FlowView").SetOrderBy("UpdateTime").SetLimt(CurrentPage, PageSize).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSqlCount, err := sqldb.SetTabName("FlowView").GetCountSql()
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
		var temp model.FlowView
		temp.ID = util.ToString(data[i]["ID"])
		temp.No = util.ToString(data[i]["No"])
		temp.StepNum = util.ToInt(data[i]["StepNum"])
		temp.Status = util.ToInt(data[i]["Status"])
		temp.MenuId = util.ToString(data[i]["MenuId"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.Notes = util.ToString(data[i]["Notes"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.MenuCode = util.ToString(data[i]["MenuCode"])
		temp.MenuName = util.ToString(data[i]["MenuName"])
		list = append(list, temp)
	}
	return &list, num, nil
}
func GetFlowByID(ID string) (*model.FlowBillModel, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case string:
				err = errors.New("新增数据异常:" + x)
			case error:
				err = errors.New("新增数据异常:" + x.Error())
			default:
				err = errors.New("新增数据异常")
			}
			if err != nil {
				fmt.Println("recover后的错误:", err)
			}

		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data model.FlowBillModel
	var main model.FlowView
	strSql := fmt.Sprintf("select * from FlowView where  ID='%s'", ID)
	dataMain, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	if len(dataMain) <= 0 {
		return nil, errors.New("未找到对应单据")
	}
	if len(dataMain) > 1 {
		return nil, errors.New("找到多个对应单据")
	}
	main.ID = util.ToString(dataMain[0]["ID"])
	main.No = util.ToString(dataMain[0]["No"])
	main.MenuId = util.ToString(dataMain[0]["MenuId"])
	main.Description = util.ToString(dataMain[0]["Description"])
	main.Notes = util.ToString(dataMain[0]["Notes"])
	main.CreateBy = util.ToString(dataMain[0]["CreateBy"])
	main.UpdateBy = util.ToString(dataMain[0]["UpdateBy"])
	main.Status = util.ToInt(dataMain[0]["Status"])
	main.StepNum = util.ToInt(dataMain[0]["StepNum"])
	main.MenuCode = util.ToString(dataMain[0]["MenuCode"])
	main.MenuName = util.ToString(dataMain[0]["MenuName"])
	createTime, _ := util.AnyToTimeStr(dataMain[0]["CreateTime"])
	updateTime, _ := util.AnyToTimeStr(dataMain[0]["UpdateTime"])
	main.CreateTime = createTime
	main.UpdateTime = updateTime
	flowStep := make([]model.FlowStep, 0)
	strSql = fmt.Sprintf("select * from FlowStep where  FlowId='%s'  order by StepNum ASC ", ID)
	dataFlowStep, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	if len(dataFlowStep) > 1 {
		for i := 0; i < len(dataFlowStep); i++ {
			var flowStepItem model.FlowStep
			flowStepItem.ID = util.ToString(dataFlowStep[i]["ID"])
			flowStepItem.FlowId = util.ToString(dataFlowStep[i]["FlowId"])
			flowStepItem.StepNum = util.ToInt(dataFlowStep[i]["StepNum"])
			flowStepItem.StepType = util.ToInt(dataFlowStep[i]["StepType"])
			flowStepItem.Notes = util.ToString(dataFlowStep[i]["Notes"])
			flowStepItem.CreateBy = util.ToString(dataFlowStep[i]["CreateBy"])
			flowStepItem.UpdateBy = util.ToString(dataFlowStep[i]["UpdateBy"])
			createTime, _ = util.AnyToTimeStr(dataFlowStep[i]["CreateTime"])
			updateTime, _ = util.AnyToTimeStr(dataFlowStep[i]["UpdateTime"])
			flowStepItem.CreateTime = createTime
			flowStepItem.UpdateTime = updateTime
			flowStep = append(flowStep, flowStepItem)
		}
	}
	flowStepUser := make([]model.FlowStepUserView, 0)
	strSql = fmt.Sprintf("select * from FlowStepUserView where  FlowId='%s'", ID)
	dataFlowStepUser, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}

	if len(dataFlowStepUser) > 1 {
		for i := 0; i < len(dataFlowStepUser); i++ {
			var flowStepUserItem model.FlowStepUserView
			flowStepUserItem.ID = util.ToString(dataFlowStepUser[i]["ID"])
			flowStepUserItem.FlowId = util.ToString(dataFlowStepUser[i]["FlowId"])
			flowStepUserItem.UserId = util.ToString(dataFlowStepUser[i]["UserId"])
			flowStepUserItem.StepId = util.ToString(dataFlowStepUser[i]["StepId"])
			flowStepUserItem.AccountName = util.ToString(dataFlowStepUser[i]["AccountName"])
			flowStepUserItem.RealName = util.ToString(dataFlowStepUser[i]["RealName"])
			flowStepUserItem.Notes = util.ToString(dataFlowStepUser[i]["Notes"])
			flowStepUserItem.CreateBy = util.ToString(dataFlowStepUser[i]["CreateBy"])
			flowStepUserItem.UpdateBy = util.ToString(dataFlowStepUser[i]["UpdateBy"])
			createTime, _ = util.AnyToTimeStr(dataFlowStepUser[i]["CreateTime"])
			updateTime, _ = util.AnyToTimeStr(dataFlowStepUser[i]["UpdateTime"])
			flowStepUserItem.CreateTime = createTime
			flowStepUserItem.UpdateTime = updateTime
			flowStepUser = append(flowStepUser, flowStepUserItem)
		}
	}
	data.Flow = main
	data.FlowStep = flowStep
	data.FlowStepUser = flowStepUser
	// util.StructCopy(&((*data).Flow), &(main[0]))
	// util.StructCopy(&((*data).FlowStep), &flowStep)
	// util.StructCopy(&((*data).FlowStepUser), &flowStepUser)
	return &data, nil
}
func GetFlowInfoByMenuId(ID string) (*model.FlowView, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.FlowView
	err = db.Table(&data).Where("MenuId", "=", ID).Where("Status", ">=", 5).Select()
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
	return &data[0], nil
}

func AddFlow(data model.FlowBillModel) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case string:
				err = errors.New("新增数据异常:" + x)
			case error:
				err = errors.New("新增数据异常:" + x.Error())
			default:
				err = errors.New("新增数据异常")
			}
			if err != nil {
				fmt.Println("recover后的错误:", err)
			}

		}
	}()
	timeStr := util.GetNowStr()
	data.Flow.CreateTime = timeStr
	data.Flow.UpdateTime = timeStr
	main := new(model.Flow)
	util.StructCopy(main, &data.Flow)
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	billNo, err := GetBillNo("LC")
	if err != nil {
		return err
	}
	(*main).No = billNo
	db.Begin()
	_, err = db.Table("Flow").ExtraCols(consts.GetFlowInfo()...).Insert(main)
	if err != nil {
		db.Rollback()
		return err
	}
	for i := 0; i < len(data.FlowStep); i++ {
		item := new(model.FlowStep)
		util.StructCopy(item, &(data.FlowStep[i]))
		(*item).CreateBy = (*main).CreateBy
		(*item).UpdateBy = (*main).UpdateBy
		(*item).CreateTime = timeStr
		(*item).UpdateTime = timeStr
		_, err = db.Table("FlowStep").ExtraCols(consts.GetFlowInfo()...).Insert(item)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	for i := 0; i < len(data.FlowStepUser); i++ {
		item := new(model.FlowStepUser)
		util.StructCopy(item, &(data.FlowStepUser[i]))
		(*item).CreateBy = (*main).CreateBy
		(*item).UpdateBy = (*main).UpdateBy
		(*item).CreateTime = timeStr
		(*item).UpdateTime = timeStr

		_, err = db.Table("FlowStepUser").ExtraCols(consts.GetFlowInfo()...).Insert(item)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err

}
func UpdateFlow(data model.FlowBillModel) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case string:
				err = errors.New("新增数据异常:" + x)
			case error:
				err = errors.New("新增数据异常:" + x.Error())
			default:
				err = errors.New("新增数据异常")
			}
			if err != nil {
				fmt.Println("recover后的错误:", err)
			}

		}
	}()
	timeStr := util.GetNowStr()
	data.Flow.UpdateTime = timeStr
	main := new(model.Flow)
	util.StructCopy(main, &data.Flow)
	db, err := util.OpenDB()
	delSql1 := ""
	delSql2 := ""
	if err != nil {
		return err
	}
	//删除FlowStep SQL
	strSql1 := fmt.Sprintf("select ID from FlowStep where FlowId='%s' ", data.Flow.ID)
	itemData1, err := db.Query(strSql1)
	if err != nil {
		return err
	}
	delList1 := make([]string, 0)
	for i := 0; i < len(itemData1); i++ {
		findInfo := false
		for j := 0; j < len(data.FlowStep); j++ {
			if util.ToString(itemData1[i]["ID"]) == data.FlowStep[j].ID {
				findInfo = true
				break
			}
		}
		if !findInfo {
			delList1 = append(delList1, "'"+util.ToString(itemData1[i]["ID"])+"'")
		}
	}
	if len(delList1) > 0 {
		delwhereSql := strings.Join(delList1, ",")
		delSql1 = fmt.Sprintf("Delete from FlowStep where ID in(%s)", delwhereSql)
	}
	//删除FlowStepUser SQL
	strSql2 := fmt.Sprintf("select ID from FlowStepUser where FlowId='%s' ", data.Flow.ID)
	itemData2, err := db.Query(strSql2)
	if err != nil {
		return err
	}
	delList2 := make([]string, 0)
	for i := 0; i < len(itemData2); i++ {
		findInfo := false
		for j := 0; j < len(data.FlowStepUser); j++ {
			if util.ToString(itemData2[i]["ID"]) == data.FlowStepUser[j].ID {
				findInfo = true
				break
			}
		}
		if !findInfo {
			delList2 = append(delList2, "'"+util.ToString(itemData2[i]["ID"])+"'")
		}
	}
	if len(delList2) > 0 {
		delwhereSql := strings.Join(delList2, ",")
		delSql2 = fmt.Sprintf("Delete from FlowStepUser where ID in(%s)", delwhereSql)
	}

	db.Begin()
	//先删除了再说
	if len(delList1) > 0 {
		_, err = db.Execute(delSql1)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	if len(delList2) > 0 {
		_, err = db.Execute(delSql2)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	//再更新
	_, err = db.Table("Flow").ExtraCols(consts.GetFlowInfo()...).Where("ID", "=", (*main).ID).Update(main)
	if err != nil {
		db.Rollback()
		return err
	}
	for i := 0; i < len(data.FlowStep); i++ {
		item := new(model.FlowStep)
		util.StructCopy(item, &(data.FlowStep[i]))
		(*item).UpdateBy = (*main).UpdateBy
		(*item).UpdateTime = timeStr
		count, err := db.Table("FlowStep").ExtraCols(consts.GetFlowInfo()...).Where("ID", "=", (*item).ID).Update(item)
		if err != nil {
			db.Rollback()
			return err
		}
		if count <= 0 {
			(*item).CreateBy = (*main).UpdateBy
			(*item).CreateTime = timeStr
			_, err = db.Table("FlowStep").ExtraCols(consts.GetFlowInfo()...).Insert(item)
			if err != nil {
				db.Rollback()
				return err
			}

		}
	}
	for i := 0; i < len(data.FlowStepUser); i++ {
		item := new(model.FlowStepUser)
		util.StructCopy(item, &(data.FlowStepUser[i]))
		(*item).UpdateBy = (*main).UpdateBy
		(*item).UpdateTime = timeStr

		count, err := db.Table("FlowStepUser").ExtraCols(consts.GetFlowInfo()...).Where("ID", "=", (*item).ID).Update(item)
		if err != nil {
			db.Rollback()
			return err
		}
		if count <= 0 {
			(*item).CreateBy = (*main).UpdateBy
			(*item).CreateTime = timeStr
			_, err = db.Table("FlowStepUser").ExtraCols(consts.GetFlowInfo()...).Insert(item)
			if err != nil {
				db.Rollback()
				return err
			}
		}
	}
	db.Commit()
	return err

}
func DeleteFlow(idList []string) error {
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
		count, err := db.Table("Flow").Where("ID", v).Where("Status", "=", 0).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		if count > 0 {
			_, err = db.Table("FlowStep").Where("FlowId", v).Delete()
			if err != nil {
				db.Rollback()
				return err
			}
			_, err = db.Table("FlowStepUser").Where("FlowId", v).Delete()
			if err != nil {
				db.Rollback()
				return err
			}
		}
	}
	db.Commit()
	return err
}
func UpdateFlowStatus(ID, UpdateBy string, oldStatus, newStatus int64) error {
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	timeStr := util.GetNowStr()
	count, err := db.Table("Flow").Data(map[string]interface{}{"Status": newStatus, "UpdateBy": UpdateBy, "UpdateTime": timeStr}).Where("ID", "=", ID).Where("Status", "=", oldStatus).Update()
	if err != nil {
		return err
	}
	if count <= 0 {
		return errors.New("数据已修改")
	}
	return err
}
