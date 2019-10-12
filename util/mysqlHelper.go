package util

import (
	"GoModDemo/model"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)
func GetWhereSqlOrderLimt(TabName,ParameterStr string,OrderStr string,SortStr string, PageSize, CurrentPage int) (string, error) {
	
	whereSql,err:=GetWhereSql(ParameterStr)
	if err != nil {
	   return "", err
	}
	 
	whereSql ="select * from  " +TabName+"  where "+ whereSql+" Order By "+OrderStr+" "+SortStr+ "  LIMIT " + strconv.Itoa( (CurrentPage-1)*PageSize)+"," +strconv.Itoa(CurrentPage*PageSize)
	return whereSql,nil
}
func GetWhereSqlLimt(TabName,ParameterStr string,PageSize, CurrentPage int) (string, error) {
	
	 whereSql,err:=GetWhereSql(ParameterStr)
	 if err != nil {
		return "", err
	 }
	  
	 whereSql ="select * from  " +TabName+"  where "+ whereSql+ "  LIMIT " + strconv.Itoa( (CurrentPage-1)*PageSize)+"," +strconv.Itoa(CurrentPage*PageSize)
	 return whereSql,nil
}
func GetWhereSqlCount(TabName,ParameterStr string) (string, error) {
	
	whereSql,err:=GetWhereSql(ParameterStr)
	if err != nil {
	   return "", err
	}
	 
	whereSql ="select count(1) as Num from  " +TabName+"  where "+ whereSql
	return whereSql,nil
}
func GetWhereSql(ParameterStr string) (string, error) {
	sqlWhere := " 1=1 "
	if ParameterStr == "" || ParameterStr == "[]" {
		return sqlWhere, nil
	}
	var filterModelList []model.FilterModel

	err := json.Unmarshal([]byte(ParameterStr), &filterModelList)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(filterModelList); i++ {
		fieldWhere,err := getFieldWhere(&(filterModelList[i]))
		if err != nil {
			return "", err
		}
		if sqlWhere == "" {
			sqlWhere = "1=1 " + fieldWhere
		} else {
			sqlWhere+= fieldWhere
		}

	}
	return sqlWhere, nil
}
func getFieldWhere(model *model.FilterModel) (string, error) {
	strTemp:=""
	if model==nil {
		return "", errors.New("对象为空")
	}
	if (*model).Logic=="" {
		return "", errors.New("关系符为空")
	}
	if (*model).Action=="" {
		return "", errors.New("算术符为空")
	}
	strTemp=" " + (*model).Logic
	strTemp+=" " +(*model).Column+ " "+(*model).Action 
	strTemp+= getwhereByDataType( (*model).DataType)
	strTemp+= getwhereByAction( (*model).Action)
	strTemp+=(*model).Value
	strTemp+= getwhereByAction( (*model).Action)
	strTemp+=getwhereByDataType( (*model).DataType)

	return strTemp,nil

}
func  getwhereByAction(action string)  string {
	action =strings.ToLower(action)
	switch(action) {
	case "like":
	 return "%"
	default:
	 return ""
}
}

func  getwhereByDataType(dataType string)  string {
		switch(dataType) {
		case "S","D":
		 return "'"
		case "I","F":
			return ""
		default:
			return "'"
	}
}
