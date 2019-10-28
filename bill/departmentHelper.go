package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"

	// "encoding/json"
	"errors"
	"fmt"
	"strings"
)

func GetAllDeptForTransfer() (*[]model.TransferModel, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()

	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var dept []model.Department
	var list []model.TransferModel
	err = db.Table(&dept).Select()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(dept); i++ {
		var temp model.TransferModel
		temp.Key = dept[i].ID
		temp.Label = dept[i].Name
		temp.Title = dept[i].Name
		temp.Disabled = false
		list = append(list, temp)
	}
	return &list, err
}

func GetDeptByUserIdForTransfer(userId string) (string, error) {

	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()

	db, err := util.OpenDB()
	if err != nil {
		return "", err
	}
	strSql := fmt.Sprintf("select ud.DepartmentId from  UserDepartment ud ,Department  d   where ud.DepartmentId=d.ID and d.IsAble=1 and ud.UserId='%s'", userId)
	data, err := db.Query(strSql)
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", nil
	}
	var list []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["DepartmentId"])
		list = append(list, temp)
	}
	str := strings.Join(list, ",")
	return str, err

}

func GetAllDeptInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.Department, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	whereSql, err := util.GetWhereSqlLimt("Department", ParameterStr, PageSize, CurrentPage)
	if err != nil {
		return nil, err
	}
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	fmt.Println(whereSql)
	data, err := db.Query(whereSql)
	if err != nil {
		return nil, err
	}
	list := make([]model.Department, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.Department
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.ParentId = util.ToString(data[i]["ParentId"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		temp.Sort = util.ToInt(data[i]["Sort"])
		list = append(list, temp)
	}
	return &list, nil
}

func DeleteDept(idList []string) error {
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
		_, err = db.Table("Department").Where("ID", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("UserDepartment").Where("DepartmentId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}

	}
	db.Commit()
	return err
}

func AddDept(data model.Department) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.CreateTime = timeStr
	data.UpdateTime = timeStr
	data.Description = ""
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.ExtraCols(consts.GetTabInfo()...).Insert(&data)
	if err != nil {
		return err
	}
	return err
}

func UpdateDept(data model.Department) error {
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
	_, err = db.ExtraCols(consts.GetTabInfo()...).Where("ID", data.ID).Update(&data)
	if err != nil {
		return err
	}
	return err
}

func GetDeptByID(ID string) (*model.Department, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.Department
	err = db.Table(&data).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("未找到对应权限")
	}
	if len(data) > 1 {
		return nil, errors.New("找到多个对应权限")
	}
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime, _ = util.ParseAnyToStr(data[0].UpdateTime)
	return &data[0], nil
}
