package bill

import (
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
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

func GetAllMenuInfo(ParameterStr string,PageSize, CurrentPage int)(*[]model.Menu, error)   {
	whereSql,err:=util.GetWhereSqlLimt("Menu" ,ParameterStr,PageSize,CurrentPage)
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
   list := make([]model.Menu, 0)
   if len(data) <= 0 {
	   return &list, nil
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
	   temp.CreateTime= createTime
	   temp.UpdateTime= updateTime
	   temp.IsAble = util.ToInt(data[i]["IsAble"])
	   temp.Sort = util.ToInt(data[i]["Sort"])
	   list = append(list, temp)
   }
   return &list,nil
}
func GetCascaderMenu ()(*[]model.CascaderMenu, error) {

	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	list := make([]model.CascaderMenu, 0)
	menuList:=make([]model.Menu, 0)
	err = db.Table(&menuList).Select()
	if err != nil {
		return nil, err
	}
	if len(menuList) <= 0 {
		return &list, nil
	}
	for i := 0; i < len(menuList); i++ {
		var temp model.CascaderMenu
		temp.Value = util.ToString(menuList[i].ID)
		temp.Label = util.ToString(menuList[i].Name)
		list = append(list, temp)
	}
	list = *generateCascaderMenu(&menuList)
	return &list, nil

}
func generateCascaderMenu(list *[]model.Menu) *[]model.CascaderMenu {
	listTemp := make([]model.CascaderMenu, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == "0" {
		    var	temp model.CascaderMenu
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
func generateCascaderMenuNext(id string, list *[]model.Menu) *[]model.CascaderMenu {
	listTemp := make([]model.CascaderMenu, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == id {
		    var	temp model.CascaderMenu
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