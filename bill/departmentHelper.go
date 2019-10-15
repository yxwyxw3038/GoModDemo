package bill

import (
	"GoModDemo/model"
	"GoModDemo/util"
	// "GoModDemo/consts"
	// "encoding/json"
	"errors"
	"fmt"
	"strings"
	
)
func GetAllDeptForTransfer()(*[]model.TransferModel, error)  {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("数据异常")
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
	for i:=0;i<len(dept);i++ {
	   var temp model.TransferModel
	   temp.Key=dept[i].ID
	   temp.Label=dept[i].Name
	   temp.Title=dept[i].Name
	   temp.Disabled=false
	   list=append(list,temp)
	}
	return &list, err
}

func GetDeptByUserIdForTransfer(userId string) (string, error)  {

	var err error
	defer func() {
		if p := recover(); p != nil {
			err=errors.New("数据异常")
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
	str:=strings.Join(list, ",")
	return str, err

}

