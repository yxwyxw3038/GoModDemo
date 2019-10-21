package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"

	// "encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func GetAllRoleForTransfer() (*[]model.TransferModel, error) {
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
	var role []model.Role
	var list []model.TransferModel
	err = db.Table(&role).Select()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(role); i++ {
		var temp model.TransferModel
		temp.Key = role[i].ID
		temp.Label = role[i].Name
		temp.Title = role[i].Name
		temp.Disabled = false
		list = append(list, temp)
	}
	return &list, err
}

func GetRoleByUserIdForTransfer(userId string) (string, error) {

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
	str := strings.Join(list, ",")
	return str, err

}

func GetAllRoleInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.Role, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	whereSql, err := util.GetWhereSqlLimt("Role", ParameterStr, PageSize, CurrentPage)
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
	list := make([]model.Role, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.Role
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		list = append(list, temp)
	}
	return &list, nil
}

func DeleteRole(idList []string) error {
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
		temp, err := util.DelSqlByField("Role", "ID", v)
		if err != nil {
			return err
		}
		sqlList = append(sqlList, temp)
	}
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	db.Begin()
	for _, v := range sqlList {
		_, err := db.Execute(v)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err
}

func AddRole(data model.Role) error {
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

func UpdateRole(data model.Role) error {
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
func GetRoleByID(ID string) (*model.Role, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.Role
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

func GetMenuByRoleIdForTree(ID string) (string, error) {
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
	strSql := fmt.Sprintf("select rm.MenuId from RoleMenu rm where  rm.RoleId='%s' and not EXISTS (select 1 from Menu  m where  rm.MenuId=m.ParentId )", ID)
	data, err := db.Query(strSql)
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", nil
	}
	var list []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["MenuId"])
		list = append(list, temp)
	}
	str := strings.Join(list, ",")
	return str, err
}

// func SetMapByBasemap(hasmap,basehasmap map[string]string){

// }
func SetMenuRole(roleId, menuStr string) error {
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
	menuList := make([]string, 0)
	if menuStr != "" {
		menuList = strings.Split(menuStr, ",")
	}
	basehasmap := make(map[string]string, 0)
	hasmap := make(map[string]string, 0)
	oldhasmap := make(map[string]string, 0)
	//加载基础资料
	strSql := "select  m.ID,m.ParentId  from Menu m "
	baseData, err := db.Query(strSql)
	if err != nil {
		return err
	}

	for i := 0; i < len(baseData); i++ {
		tempMenuId := util.ToString(baseData[i]["ID"])
		tempParentId := util.ToString(baseData[i]["ParentId"])
		basehasmap[tempMenuId] = tempParentId
	}
	//补齐ParentId
	for i := 0; i < len(menuList); i++ {
		if v, ok := basehasmap[menuList[i]]; ok {
			hasmap[menuList[i]] = v
		}
	}
	//补上级菜单
	for _, v := range hasmap {
		if v == "0" {

		} else {
			if v, ok := hasmap[v]; !ok {
				if v1, ok := basehasmap[v]; ok {
					hasmap[v] = v1
				}
			}
		}
	}
	strSql = fmt.Sprintf("select rm.MenuId,m.ParentId from RoleMenu rm ,Menu  m where  rm.RoleId='%s' and rm.MenuId=m.ID", roleId)
	data, err := db.Query(strSql)
	if err != nil {
		return err
	}
	//在本权限所有菜单清单中找  找不到就清除
	var delList []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["MenuId"])
		if _, ok := hasmap[temp]; !ok {

			delList = append(delList, temp)
		}
		oldhasmap[temp] = ""
	}

	var tempList []model.RoleMenu
	timeStr := util.GetNowStr()
	for key, _ := range hasmap {
		if _, ok := oldhasmap[key]; !ok {

			newUuid := uuid.New()
			newUuidStr := newUuid.String()
			var temp model.RoleMenu
			temp.ID = newUuidStr
			temp.MenuId = key
			temp.RoleId = roleId
			temp.CreateBy = "admin"
			temp.CreateTime = timeStr
			temp.UpdateBy = "admin"
			temp.UpdateTime = timeStr
			tempList = append(tempList, temp)
		}

	}

	db.Begin()
	for i := 0; i < len(delList); i++ {
		_, err := db.Table("RoleMenu").Where("RoleId", roleId).Where("MenuId", delList[i]).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("RoleMenuButton").Where("RoleId", roleId).Where("MenuId", delList[i]).Delete()
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

func SetButtonByMenuIdRoleId(menuId, roleId, buttonStr string) error {
	var err error
	hasmap := make(map[string]string, 0)
	oldhasmap := make(map[string]string, 0)
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	buttonList := make([]string, 0)
	if buttonStr != "" {
		buttonList = strings.Split(buttonStr, ",")
	}
	for i := 0; i < len(buttonList); i++ {
		hasmap[buttonList[i]] = ""
	}
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	strSql := fmt.Sprintf("select r.ButtonId from RoleMenuButton r  where r.MenuId='%s' and r.RoleId='%s' and EXISTS (select  1 from MenuButton as m where m.ButtonId=r.ButtonId and m.MenuId=r.MenuId )", menuId, roleId)
	data, err := db.Query(strSql)
	if err != nil {
		return err
	}
	var delList []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["ButtonId"])
		if _, ok := hasmap[temp]; !ok {

			delList = append(delList, temp)
		}
		oldhasmap[temp] = ""
	}

	var tempList []model.RoleMenuButton
	timeStr := util.GetNowStr()
	for key, _ := range hasmap {
		if _, ok := oldhasmap[key]; !ok {

			newUuid := uuid.New()
			newUuidStr := newUuid.String()
			var temp model.RoleMenuButton
			temp.ID = newUuidStr
			temp.MenuId = menuId
			temp.RoleId = roleId
			temp.ButtonId = key
			temp.CreateBy = "admin"
			temp.CreateTime = timeStr
			temp.UpdateBy = "admin"
			temp.UpdateTime = timeStr

			tempList = append(tempList, temp)
		}

	}

	db.Begin()
	for i := 0; i < len(delList); i++ {
		_, err := db.Table("RoleMenuButton").Where("MenuId", menuId).Where("RoleId", roleId).Where("ButtonId", delList[i]).Delete()
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
