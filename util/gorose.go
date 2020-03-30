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
	// dbstr := setting.DatabaseSetting.User + ":" + setting.DatabaseSetting.Password + "@tcp(" + setting.DatabaseSetting.Host + ")/" + setting.DatabaseSetting.Name + "?charset=utf8&parseTime=true"
	// config0 := gorose.Config{Driver: setting.DatabaseSetting.Type, Dsn: dbstr}

	// dbstr1 := setting.DatabaseSetting1.User + ":" + setting.DatabaseSetting1.Password + "@tcp(" + setting.DatabaseSetting1.Host + ")/" + setting.DatabaseSetting1.Name + "?charset=utf8&parseTime=true"
	// config1 := gorose.Config{Driver: setting.DatabaseSetting1.Type, Dsn: dbstr1}

	// var configCluster = &gorose.ConfigCluster{
	// 	Master: []gorose.Config{config0},
	// 	Slave:  []gorose.Config{config1},
	// 	Driver: setting.DatabaseSetting.Type,
	// }
	// engin, err := gorose.Open(configCluster)
	if err != nil {
		return nil, err
	}
	return engin.NewOrm(), nil
}

func ExecuteList(orm gorose.IOrm, sql ...string) error {
	orm.Begin()
	for _, v := range sql {
		_, err := orm.Execute(v)
		if err != nil {
			orm.Rollback()
			return err
		}
	}
	orm.Commit()
	return nil
}

// func ExecuteListObj(orm gorose.IOrm,opt string, obj ...interface{})  error{
// 	orm.Begin()
//     var err  error
// 	for _,v :=range obj{
// 		switch opt {
// 		case "Insert":
// 			_,err = orm.Insert(&v)
// 		default:
// 			_,err = orm.Update(&v)
// 		}

// 		if (err!=nil) {
// 			orm.Rollback()
// 			return err
// 		}
// 	}
// 	orm.Commit()
// 	return nil
// }
