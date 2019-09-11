package bill
import (
	"GoModDemo/model"
	"GoModDemo/util"
	"errors"
)
func UserAuth(userName string,passWord string) error {
	
	db, err := util.OpenDB()
	if err != nil {
		return err
	} 
	var user []model.User
	err = db.Table(&user).Where("AccountName", "=", userName).Select()
    if err != nil {
		return  err
	} 
	if len(user)<=0 {
		return  errors.New("未找到对应用户")
	}
	if len(user)>1 {
		return errors.New("找到多个对应用户")
	}
	if user[0].PassWord!=passWord {
		return errors.New("帐户与密码不一至")
	}
	return nil
}

func GetUserInfoByAccountName(userName string) (*model.User, error) {
	
	db, err := util.OpenDB()
	if err != nil {
		return nil,err
	} 
	var user []model.User
	err = db.Table(&user).Where("AccountName", "=", userName).Select()
    if err != nil {
		return nil, err
	} 
	if len(user)<=0 {
		return nil, errors.New("未找到对应用户")
	}
	if len(user)>1 {
		return nil, errors.New("找到多个对应用户")
	}
	
	return &user[0],nil
}