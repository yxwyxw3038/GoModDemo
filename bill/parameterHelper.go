package bill

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
	"fmt"
	// "github.com/google/uuid"
)

func GetTreeParameter() (*[]model.TreeModel, error) {

	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	list := make([]model.TreeModel, 0)
	dataList := make([]model.Parameter, 0)
	err = db.Table(&dataList).Select()
	if err != nil {
		return nil, err
	}
	if len(dataList) <= 0 {
		return &list, nil
	}

	list = *generateTreeParameter(&dataList)
	return &list, nil

}
func generateTreeParameter(list *[]model.Parameter) *[]model.TreeModel {
	listTemp := make([]model.TreeModel, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == "0" {
			var temp model.TreeModel
			temp.ID = util.ToString((*list)[i].ID)
			temp.Label = util.ToString((*list)[i].Name)
			node := generateTreeParameterNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Children = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}
func generateTreeParameterNext(id string, list *[]model.Parameter) *[]model.TreeModel {
	listTemp := make([]model.TreeModel, 0)
	for i := 0; i < len(*list); i++ {
		if (*list)[i].ParentId == id {
			var temp model.TreeModel
			temp.ID = util.ToString((*list)[i].ID)
			temp.Label = util.ToString((*list)[i].Name)
			node := generateTreeParameterNext((*list)[i].ID, list)
			if node != nil && len(*node) > 0 {
				temp.Children = *node
			}
			listTemp = append(listTemp, temp)
		}
	}
	return &listTemp
}
func GetAllParameterInfo(ParameterStr string, PageSize, CurrentPage int) (*[]model.Parameter, int, error) {
	list := make([]model.Parameter, 0)
	whereSql, err := util.GetWhereSqlOrderLimt("Parameter", ParameterStr, "Sort", consts.ASC, PageSize, CurrentPage)
	if err != nil {
		return nil, 0, err
	}
	whereSqlCount, err := util.GetWhereSqlCount("Parameter", ParameterStr)
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
		var temp model.Parameter
		temp.ID = util.ToString(data[i]["ID"])
		temp.Name = util.ToString(data[i]["Name"])
		temp.ParentId = util.ToString(data[i]["ParentId"])
		temp.Code = util.ToString(data[i]["Code"])
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

func GetParameterByID(ID string) (*model.Parameter, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	var data []model.Parameter
	err = db.Table(&data).Where("ID", "=", ID).Select()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("未找到对应参数")
	}
	if len(data) > 1 {
		return nil, errors.New("找到多个对应参数")
	}
	data[0].CreateTime, _ = util.ParseAnyToStr(data[0].CreateTime)
	data[0].UpdateTime, _ = util.ParseAnyToStr(data[0].UpdateTime)
	return &data[0], nil
}
func DeleteParameter(idList []string) error {
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
		_, err = db.Table("Parameter").Where("ID", v).Delete()
		if err != nil {
			db.Rollback()
			return err
		}

	}
	db.Commit()
	return err
}

func AddParameter(data model.Parameter) error {
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
	_, err = db.ExtraCols(consts.GetTabInfo()...).Insert(&data)
	if err != nil {
		return err
	}
	return err
}

func UpdateParameter(data model.Parameter) error {
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
