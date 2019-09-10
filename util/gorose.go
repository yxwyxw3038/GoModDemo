package util

import (
	"GoModDemo/setting"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose/v2"
)

func OpenDB() (gorose.IOrm, error) {
	var engin *gorose.Engin
	dbstr := setting.DatabaseSetting.User + ":" + setting.DatabaseSetting.Password + "@tcp(" + setting.DatabaseSetting.Host + ")/" + setting.DatabaseSetting.Name + "?charset=utf8&parseTime=true"
	engin, err := gorose.Open(&gorose.Config{Driver: "mysql", Dsn: dbstr})
	if err != nil {
		return nil, err
	}
	return engin.NewOrm(), nil
}
