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
func GetAllRoleForTransfer()(*[]model.TransferModel, error)  {
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
	var role []model.Role
	var list []model.TransferModel
	err = db.Table(&role).Select()
	if err != nil {
		return nil, err
	}
	for i:=0;i<len(role);i++ {
	   var temp model.TransferModel
	   temp.Key=role[i].ID
	   temp.Label=role[i].Name
	   temp.Title=role[i].Name
	   temp.Disabled=false
	   list=append(list,temp)
	}
	return &list, err
}

func GetRoleByUserIdForTransfer(userId string) (string, error)  {

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
	strSql := fmt.Sprintf("select ud.RoleId from  UserRole ud ,Role   d   where ud.RoleId=d.ID and d.IsAble=1 and ud.UserId='%s'", userId)
	data, err := db.Query(strSql)
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", nil
	}
	var list []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["RoleId"])
		list = append(list, temp)
	}
	str:=strings.Join(list, ",")
	return str, err

}