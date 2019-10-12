package bill

import (
	"GoModDemo/model"
	"GoModDemo/util"
	"fmt"
)
func GetButtonByMenuIdAndUserId(menuId string,userId string)(*[]model.Button, error) {
	db, err := util.OpenDB()
	if err != nil {
		return nil, err
	}
	strSql := fmt.Sprintf("select b.* from Button  b ,RoleMenuButton rmb,UserRole ur where b.IsAble=1 and rmb.ButtonId=b.ID and rmb.RoleId=ur.RoleId  and rmb.MenuId='%s' and ur.UserId='%s' ORDER BY b.Sort",menuId, userId)
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
		temp.CreateTime= createTime
		temp.UpdateTime= updateTime
		temp.Sort = util.ToInt(data[i]["Sort"])
		temp.IsAble = util.ToInt(data[i]["IsAble"])
		list = append(list, temp)
	}
	return &list, nil

}