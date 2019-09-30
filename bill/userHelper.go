package bill

import (
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
)

func UserAuth(userName string, passWord string) error {

	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	var user []model.User
	err = db.Table(&user).Where("AccountName", "=", userName).Select()
	if err != nil {
		return err
	}
	if len(user) <= 0 {
		return errors.New("未找到对应用户")
	}
	if len(user) > 1 {
		return errors.New("找到多个对应用户")
	}
	if user[0].PassWord != passWord {
		return errors.New("帐户与密码不一至")
	}
	return nil
}

func GetUserInfoByAccountName(userName string) (*model.User, error) {

	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var user []model.User
	err = db.Table(&user).Where("AccountName", "=", userName).Select()
	if err != nil {
		return nil, err
	}
	if len(user) <= 0 {
		return nil, errors.New("未找到对应用户")
	}
	if len(user) > 1 {
		return nil, errors.New("找到多个对应用户")
	}

	return &user[0], nil
}
func GetUserByID(ID string) (*model.User, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var user []model.User
	err = db.Table(&user).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(user) <= 0 {
		return nil, errors.New("未找到对应用户")
	}
	if len(user) > 1 {
		return nil, errors.New("找到多个对应用户")
	}
	user[0].CreateTime, _ = util.ParseAnyToStr(user[0].CreateTime)
	user[0].UpdateTime, _ = util.ParseAnyToStr(user[0].UpdateTime)
	return &user[0], nil
}
func GetUserMenu(userId string) (*[]model.MenuTree, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	strSql := fmt.Sprintf("SELECT	m.* FROM	Menu m,	RoleMenu rm,	UserRole ur,	Role r,	User u WHERE	m.IsAble != 0 AND rm.MenuId = m.ID AND ur.RoleId = rm.RoleId AND r.IsAble != 0 AND ur.UserId = u.ID AND ur.RoleId = r.ID AND u.IsAble != 0 AND u.ID = '%s'", userId)
	data, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	list := make([]model.MenuTree, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.MenuTree
		temp.ID = util.ToString(data[i]["ID"])
		temp.Icon = util.ToString(data[i]["Icon"])
		temp.MenuName = util.ToString(data[i]["Name"])
		temp.ParentId = util.ToString(data[i]["ParentId"])
		temp.Url = util.ToString(data[i]["LinkAddress"])
		list = append(list, temp)
	}
	list = *generateMenuTree(&list)
	return &list, nil
}
func generateMenuTree(list *[]model.MenuTree) *[]model.MenuTree {
	listTemp := make([]model.MenuTree, 0)

	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == "0" {
			temp := (*list)[i]
			node := generateMenuTreeNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Node = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}
func generateMenuTreeNext(id string, list *[]model.MenuTree) *[]model.MenuTree {
	listTemp := make([]model.MenuTree, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == id {
			temp := (*list)[i]
			node := generateMenuTreeNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Node = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}




func  GetAllUserInfo(ParameterStr string,PageSize, CurrentPage int)(*[]model.User, error)   {
	 whereSql,err:=util.GetWhereSqlLimt("User" ,ParameterStr,PageSize,CurrentPage)
	 db, err := util.OpenDB()
	 if err != nil {
		 return nil, err
	 }
	 fmt.Println(whereSql)
	data, err := db.Query(whereSql)
	if err != nil {
		return nil, err
	}
	list := make([]model.User, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.User
		temp.ID = util.ToString(data[i]["ID"])
		temp.AccountName = util.ToString(data[i]["AccountName"])
		temp.PassWord = util.ToString(data[i]["PassWord"])
		temp.RealName = util.ToString(data[i]["RealName"])
		temp.MobilePhone = util.ToString(data[i]["MobilePhone"])
		temp.Email = util.ToString(data[i]["Email"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime= createTime
		temp.UpdateTime= updateTime
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		temp.IfChangePwd = util.ToInt(data[i]["IfChangePwd"])
		list = append(list, temp)
	}
	return &list,nil
}