package authentication

import (

"GoModDemo/util"
"GoModDemo/bill"
)

type Auth struct {
    Username string
    Password string
}

func (a *Auth) Check() (bool, error) {
    userName := a.Username
    passWord := a.Password
    newpassword:=util.Md5(passWord)
    // todo：实现自己的鉴权逻辑
    err :=bill.UserAuth(userName,newpassword)
    if err!=nil {
      return false, err
    }
    return true, nil
}