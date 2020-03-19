package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	"strings"

	"github.com/yxwyxw3038/whysql"
)

func GetButtonByMenuIdAndUserId(menuId string, userId string) (*[]model.Button, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	strSql := fmt.Sprintf("select b.* from Button  b ,RoleMenuButton rmb,UserRole ur where b.IsAble=1 and rmb.ButtonId=b.ID and rmb.RoleId=ur.RoleId  and rmb.MenuId='%s' and ur.UserId='%s' and EXISTS ( select 1 from MenuButton mb where mb.ButtonId=rmb.ButtonId and mb.MenuId=rmb.MenuId   ) ORDER BY b.Sort", menuId, userId)
	data, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	list := make([]model.Button, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.Button
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.Code = util.ToString(data[i]["Code"])
		temp.Icon = util.ToString(data[i]["Icon"])
		temp.Description = util.ToString(data[i]["Description"])
		temp.CreateBy = util.ToString(data[i]["CreateBy"])
		temp.UpdateBy = util.ToString(data[i]["UpdateBy"])
		createTime, _ := util.AnyToTimeStr(data[i]["CreateTime"])
		updateTime, _ := util.AnyToTimeStr(data[i]["UpdateTime"])
		temp.CreateTime = createTime
		temp.UpdateTime = updateTime
		temp.Sort = util.ToInt(data[i]["Sort"])
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		list = append(list, temp)
	}
	return &list, nil

}
func GetAllButtonForTransfer() (*[]model.TransferModel, error) {
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
	var btn []model.Button
	var list []model.TransferModel
	err = db.Table(&btn).Select()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(btn); i++ {
		var temp model.TransferModel
		temp.Key = btn[i].ID
		temp.Label = btn[i].Name
		temp.Title = btn[i].Name
		temp.Disabled = false
		list = append(list, temp)
	}
	return &list, err
}
func GetButtonByMenuIdnForTransfer(menuId string) (string, error) {

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
	strSql := fmt.Sprintf("select b.ID from Button b ,MenuButton mb where mb.ButtonId=b.ID and b.IsAble=1 and mb.MenuId='%s'", menuId)
	data, err := db.Query(strSql)
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", nil
	}
	var list []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["ID"])
		list = append(list, temp)
	}
	str := strings.Join(list, ",")
	return str, err

}
func GetAllButtonByMenuIdForTransfer(menuId string) (*[]model.TransferModel, error) {
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
	strSql := fmt.Sprintf("select b.ID,b.Name from Button b ,MenuButton mb where mb.ButtonId=b.ID and b.IsAble=1 and mb.MenuId='%s'", menuId)
	data, err := db.Query(strSql)
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, err
	}
	var list []model.TransferModel
	for i := 0; i < len(data); i++ {
		var temp model.TransferModel
		temp.Key = util.ToString(data[i]["ID"])
		temp.Label = util.ToString(data[i]["Name"])
		temp.Title = util.ToString(data[i]["Name"])
		temp.Disabled = false
		list = append(list, temp)
	}
	return &list, err
}

func GetButtonByMenuIdRoleIdForTransfer(menuId, roleId string) (string, error) {
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
	strSql := fmt.Sprintf("select r.ButtonId from RoleMenuButton as r where r.MenuId='%s' and r.RoleId='%s' and EXISTS (select  1 from MenuButton as m where r.ButtonId=m.ButtonId and r.MenuId=m.MenuId)", menuId, roleId)
	data, err := db.Query(strSql)
	if err != nil {
		return "", err
	}
	if len(data) <= 0 {
		return "", nil
	}
	var list []string
	for i := 0; i < len(data); i++ {
		temp := util.ToString(data[i]["ButtonId"])
		list = append(list, temp)
	}
	str := strings.Join(list, ",")
	return str, err
}

func GetAllButtonInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.Button, error) {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("数据异常")
		}
	}()
	// whereSql, err := util.GetWhereSqlOrderLimt("Button", ParameterStr, "Sort", consts.ASC, PageSize, CurrentPage)
	// if err != nil {
	// 	return nil, err
	// }
	sqldb, err := whysql.NewWhy(ParameterStr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	whereSql, err := sqldb.SetTabName("Button").SetOrderBy("Sort", whysql.ASC).SetLimt(CurrentPage, PageSize).GetQuerySql()
	if err != nil {
		fmt.Println(err.Error())
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
	list := make([]model.Button, 0)
	if len(data) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(data); i++ {
		var temp model.Button
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.Code = util.ToString(data[i]["Code"])
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
	return &list, nil
}

func DeleteButton(idList []string) error {
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
		_, err = db.Table("Button").Where("ID", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}

		_, err = db.Table("MenuButton").Where("ButtonId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}
		_, err = db.Table("RoleMenuButton").Where("ButtonId", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}

	}
	db.Commit()
	return err
}

func AddButton(data model.Button) error {
	var err error
	defer func() {
		if p := recover(); p != nil {
			err = errors.New("新增数据异常")
		}
	}()
	timeStr := util.GetNowStr()
	data.CreateTime = timeStr
	data.UpdateTime = timeStr
	// data.Description = ""
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

func UpdateButton(data model.Button) error {
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

func GetButtonByID(ID string) (*model.Button, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.Button
	err = db.Table(&data).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("未找到对应数据")
	}
	if len(data) > 1 {
		return nil, errors.New("找到多个对应数据")
	}
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime, _ = util.ParseAnyToStr(data[0].UpdateTime)
	return &data[0], nil
}
func GetButtonAllCount() (int, error) {
	// whereSqlCount, err := util.GetWhereSqlCount("Button", "")
	// if err != nil {
	// 	return 0, err
	// }
	sqldb, err := whysql.NewWhy("")
	if err != nil {
		fmt.Println(err.Error())
		return 0, err

	}
	whereSqlCount, err := sqldb.SetTabName("Button").GetCountSql()
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
