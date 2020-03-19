package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yxwyxw3038/whysql"
)

func GetMenuByID(ID string) (*model.Menu, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var menu []model.Menu
	err = db.Table(&menu).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(menu) <= 0 {
		return nil, errors.New("未找到对应菜单")
	}
	if len(menu) > 1 {
		return nil, errors.New("找到多个对应菜单")
	}
	menu[0].CreateTime, _ = util.ParseAnyToStr(menu[0].CreateTime)
	menu[0].UpdateTime, _ = util.ParseAnyToStr(menu[0].UpdateTime)
	return &menu[0], nil
}

func GetAllMenuInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.Menu, int, error) {
	list := make([]model.Menu, 0)
	// whereSql, err := util.GetWhereSqlOrderLimt("Menu", ParameterStr, "Sort", consts.ASC, PageSize, CurrentPage)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// whereSqlCount, err := util.GetWhereSqlCount("Menu", ParameterStr)
	// if err != nil {
	// 	return nil, 0, err
	// }
	sqldb, err := whysql.NewWhy(ParameterStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSql, err := sqldb.SetTabName("Menu").SetOrderBy("Sort", whysql.ASC).SetLimt(CurrentPage, PageSize).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSqlCount, err := sqldb.SetTabName("Menu").GetCountSql()
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
		var temp model.Menu
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.ParentId = util.ToString(data[i]["ParentId"])
		temp.Code = util.ToString(data[i]["Code"])
		temp.LinkAddress = util.ToString(data[i]["LinkAddress"])
		temp.Icon = util.ToString(data[i]["Icon"])
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
	return &list, num, nil
}

func GetAllMenuViewInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.MenuView, int, error) {
	list := make([]model.MenuView, 0)
	// whereSql, err := util.GetWhereSqlOrderLimt("MenuView", ParameterStr, "Sort", consts.ASC, PageSize, CurrentPage)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// whereSqlCount, err := util.GetWhereSqlCount("Menu", ParameterStr)
	// if err != nil {
	// 	return nil, 0, err
	// }
	sqldb, err := whysql.NewWhy(ParameterStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSql, err := sqldb.SetTabName("MenuView").SetOrderBy("Sort", whysql.ASC).SetLimt(CurrentPage, PageSize).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	whereSqlCount, err := sqldb.SetTabName("Menu").GetCountSql()
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
		var temp model.MenuView
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.ParentId = util.ToString(data[i]["ParentId"])
		temp.ParentName = util.ToString(data[i]["ParentName"])
		temp.Code = util.ToString(data[i]["Code"])
		temp.LinkAddress = util.ToString(data[i]["LinkAddress"])
		temp.Icon = util.ToString(data[i]["Icon"])
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
	return &list, num, nil
}

func GetCascaderMenu() (*[]model.CascaderModel, error) {

	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	list := make([]model.CascaderModel, 0)
	menuList := make([]model.Menu, 0)
	err = db.Table(&menuList).Select()
	if err != nil {
		return nil, err
	}
	if len(menuList) <= 0 {
		return &list, nil
	}
	// for i := 0; i < len(menuList); i++ {
	// 	var temp model.CascaderModel
	// 	temp.Value = util.ToString(menuList[i].ID)
	// 	temp.Label = util.ToString(menuList[i].Name)
	// 	list = append(list, temp)
	// }
	list = *generateCascaderMenu(&menuList)
	return &list, nil

}
func generateCascaderMenu(list *[]model.Menu) *[]model.CascaderModel {
	listTemp := make([]model.CascaderModel, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == "0" {
			var temp model.CascaderModel
			temp.Value = util.ToString((*list)[i].ID)
			temp.Label = util.ToString((*list)[i].Name)
			node := generateCascaderMenuNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Children = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}
func generateCascaderMenuNext(id string, list *[]model.Menu) *[]model.CascaderModel {
	listTemp := make([]model.CascaderModel, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == id {
			var temp model.CascaderModel
			temp.Value = util.ToString((*list)[i].ID)
			temp.Label = util.ToString((*list)[i].Name)
			node := generateCascaderMenuNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Children = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}
func GetMenuAllCount() (int, error) {
	// whereSqlCount, err := util.GetWhereSqlCount("Menu", "")
	// if err != nil {
	// 	return 0, err
	// }
	sqldb, err := whysql.NewWhy("")
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	whereSqlCount, err := sqldb.SetTabName("Menu").GetCountSql()
	if err != nil {
		return 0, err
	}
	fmt.Println(whereSqlCount)
	db, err := util.OpenDB()
	if err != nil {
		return 0, err
	}

	dataCount, err := db.Query(whereSqlCount)
	if err != nil {
		return 0, err
	}
	if len(dataCount) <= 0 {
		return 0, nil
	}
	num := util.ToInt(dataCount[0]["Num"])
	return num, nil
}

func DeleteMenu(idList []string) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("删除数据异常")
		}
	}()
	if len(idList) <= 0 {
		return nil
	}
	// var sqlList []string
	// for _, v := range idList {
	// 	temp, err := util.DelSqlByField("Menu", "ID", v)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	sqlList = append(sqlList, temp)
	// }
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	db.Begin()
	for _, v := range idList {
		_, err = db.Table("Menu").Where("ID", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("MenuButton").Where("MenuId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("RoleMenu").Where("MenuId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("RoleMenuButton").Where("MenuId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
	}
	db.Commit()
	return err
}

func AddMenu(data model.Menu) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.CreateTime = timeStr
	data.UpdateTime = timeStr
	data.Description = " "
	db, err := util.OpenDB()
	if err != nil {
		return err
	}
	_, err = db.ExtraCols(consts.GetMenuTabInfo()...).Insert(&data)
	if err != nil {
		return err
	}
	return err
}

func UpdateMenu(data model.Menu) error {
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
	_, err = db.ExtraCols(consts.GetMenuTabInfo()...).Where("ID", data.ID).Update(&data)
	if err != nil {
		return err
	}
	return err
}
func SetMenuButton(menuId, buttonStr string) error {
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
	buttonList := make([]string, 0)
	if buttonStr != "" {
		buttonList = strings.Split(buttonStr, ",")
	}
	hasmap := make(map[string]int, 0)
	oldhasmap := make(map[string]int, 0)
	for i := 0; i < len(buttonList); i++ {
		hasmap[buttonList[i]] = 1
	}
	strSql := fmt.Sprintf("select mb.ButtonId from MenuButton mb  where mb.MenuId='%s'  and EXISTS (select 1 from Button b where mb.ButtonId=b.ID and b.IsAble=1 ) ", menuId)
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
		oldhasmap[temp] = 0
	}

	var tempList []model.MenuButton
	timeStr := util.GetNowStr()
	for key, _ := range hasmap {
		if _, ok := oldhasmap[key]; !ok {

			newUuid := uuid.New()
			newUuidStr := newUuid.String()
			var temp model.MenuButton
			temp.ID = newUuidStr
			temp.MenuId = menuId
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
		_, err := db.Table("MenuButton").Where("MenuId", menuId).Where("ButtonId", delList[i]).Delete()
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
