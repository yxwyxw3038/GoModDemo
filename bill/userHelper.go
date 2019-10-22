package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func UserAuth(userName string, passWord string) error {

	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	var user []model.User
	isRedis := false
	s, err := util.GetRedisHasString("TaskUserInfoByAccountName", userName)
	if err == nil {
		isRedis = true
	}
	if isRedis {
		var tempUser model.User
		err = json.Unmarshal([]byte(s), &tempUser)
		if err != nil {
			isRedis = false
		}
		user = append(user, tempUser)
	}
	if !isRedis {

		err = db.Table(&user).Where("AccountName", "=", userName).Select()
		if err != nil {
			return err
		}
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
	strSql := fmt.Sprintf("SELECT	m.* FROM	Menu m,	RoleMenu rm,	UserRole ur,	Role r,	User u WHERE	m.IsAble != 0 AND rm.MenuId = m.ID AND ur.RoleId = rm.RoleId AND r.IsAble != 0 AND ur.UserId = u.ID AND ur.RoleId = r.ID AND u.IsAble != 0 AND u.ID = '%s' order by m.Sort ", userId)
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

func GetAllUserInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.User, error) {
	whereSql, err := util.GetWhereSqlLimt("User", ParameterStr, PageSize, CurrentPage)
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
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		temp.IfChangePwd = util.ToInt(data[i]["IfChangePwd"])
		list = append(list, temp)
	}
	return &list, nil
}

func GetAllUserViewInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.UserView, error) {
	whereSql, err := util.GetWhereSqlLimt("UserView", ParameterStr, PageSize, CurrentPage)
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
	list := make([]model.UserView, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.UserView
		temp.ID = util.ToString(data[i]["ID"])
		temp.AccountName = util.ToString(data[i]["AccountName"])
		temp.PassWord = "这不是密码"
		temp.RealName = util.ToString(data[i]["RealName"])
		temp.MobilePhone = util.ToString(data[i]["MobilePhone"])
		temp.Email = util.ToString(data[i]["Email"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		temp.IfChangePwd = util.ToInt(data[i]["IfChangePwd"])
		temp.DepartmentName = util.ToString(data[i]["DepartmentName"])
		temp.RoleName = util.ToString(data[i]["RoleName"])
		list = append(list, temp)
	}
	return &list, nil
}

func DeleteUser(idList []string) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("删除数据异常")
		}
	}()
	if len(idList) <= 0 {
		return nil
	}
	var sqlList []string
	for _, v := range idList {
		temp, err := util.DelSqlByField("User", "ID", v)
		if err != nil {
			return err
		}
		sqlList = append(sqlList, temp)
	}
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	err = util.ExecuteList(db, sqlList...)
	return err
}

func AddUser(data model.User) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	count, err := db.Table("User").Where("AccountName", "=", data.AccountName).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("帐号重复不得新增")
	}
	count, err = db.Table("User").Where("Email", "=", data.Email).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("邮箱重复不得新增")
	}
	count, err = db.Table("User").Where("MobilePhone", "=", data.MobilePhone).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("手机号重复不得新增")
	}

	timeStr := util.GetNowStr()
	data.CreateTime = timeStr
	data.UpdateTime = timeStr
	_, err = db.ExtraCols(consts.GetUserTabInfo()...).Insert(&data)
	fmt.Println(db.LastSql())
	if err != nil {
		return err
	}
	return err
}

func UpdateUser(data model.User) error {
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
	_, err = db.ExtraCols(consts.GetUserTabInfo()...).Where("ID", data.ID).Update(&data)
	if err != nil {
		return err
	}
	return err
}

func SetUserDept(userId, deptStr string) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	deptList := make([]string, 0)
	if deptStr != "" {
		deptList = strings.Split(deptStr, ",")
	}
	hasmap := make(map[string]int, 0)
	oldhasmap := make(map[string]int, 0)
	for i := 0; i < len(deptList); i++ {
		hasmap[deptList[i]] = 1
	}
	strSql := fmt.Sprintf("select ud.DepartmentId from  UserDepartment ud ,Department  d   where ud.DepartmentId=d.ID  and ud.UserId='%s'", userId)
	data, err := db.Query(strSql)
	if err != nil {
		return err
	}
	var delList []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["DepartmentId"])
		if _, ok := hasmap[temp]; !ok {

			delList = append(delList, temp)
		}
		oldhasmap[temp] = 0
	}

	var tempList []model.UserDepartment
	timeStr := util.GetNowStr()
	for key, _ := range hasmap {
		if _, ok := oldhasmap[key]; !ok {

			newUuid := uuid.New()
			newUuidStr := newUuid.String()
			var temp model.UserDepartment
			temp.ID = newUuidStr
			temp.UserId = userId
			temp.DepartmentId = key
			temp.CreateBy = "admin"
			temp.CreateTime = timeStr
			temp.UpdateBy = "admin"
			temp.UpdateTime = timeStr

			tempList = append(tempList, temp)
		}

	}

	db.Begin()
	for i := 0; i < len(delList); i++ {
		_, err := db.Table("UserDepartment").Where("UserId", userId).Where("DepartmentId", delList[i]).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
	}
	for i := 0; i < len(tempList); i++ {
		_, err := db.Insert(&(tempList[i]))
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err

}

func SetUserRole(userId, roleStr string) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	roleList := make([]string, 0)
	if roleStr != "" {
		roleList = strings.Split(roleStr, ",")
	}
	hasmap := make(map[string]int, 0)
	oldhasmap := make(map[string]int, 0)
	for i := 0; i < len(roleList); i++ {
		hasmap[roleList[i]] = 1
	}
	strSql := fmt.Sprintf("select ud.RoleId from  UserRole ud ,Role  d   where ud.RoleId=d.ID  and ud.UserId='%s'", userId)
	data, err := db.Query(strSql)
	if err != nil {
		return err
	}
	var delList []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["RoleId"])
		if _, ok := hasmap[temp]; !ok {

			delList = append(delList, temp)
		}
		oldhasmap[temp] = 0
	}

	var tempList []model.UserRole
	timeStr := util.GetNowStr()
	for key, _ := range hasmap {
		if _, ok := oldhasmap[key]; !ok {

			newUuid := uuid.New()
			newUuidStr := newUuid.String()
			var temp model.UserRole
			temp.ID = newUuidStr
			temp.UserId = userId
			temp.RoleId = key
			temp.CreateBy = "admin"
			temp.CreateTime = timeStr
			temp.UpdateBy = "admin"
			temp.UpdateTime = timeStr

			tempList = append(tempList, temp)
		}

	}

	db.Begin()
	for i := 0; i < len(delList); i++ {
		_, err := db.Table("UserRole").Where("UserId", userId).Where("RoleId", delList[i]).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
	}
	for i := 0; i < len(tempList); i++ {
		_, err := db.Insert(&(tempList[i]))
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err

}
