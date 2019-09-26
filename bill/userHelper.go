package bill

import (
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
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
	// db, err := util.OpenDB()
	// if err != nil {
	// 	return nil,err
	// }
	// strSql:=fmt.Sprintf("select m.* from  UserRole ur, User u,Role r,RoleMenu rm,Menu m where   ur.UserId=u.ID and ur.RoleId=r.ID and  ur.RoleId =rm.RoleId and rm.MenuId=m.ID and m.IsAble!=0  and u.ID='%s'",userId)
	// strSql:=fmt.Sprintf("select m.* from  UserRole ur, User u,Role r,RoleMenu rm,Menu m where   ur.UserId=u.ID and ur.RoleId=r.ID and  ur.RoleId =rm.RoleId and rm.MenuId=m.ID and m.IsAble!=0  and u.ID='%s'",userId)
	// res:=db.Query(strSql)
	// fmt.Println(res)
	return nil, nil
}
