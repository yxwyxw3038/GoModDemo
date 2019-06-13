package authentication

import (
  "fmt"
"GoModDemo/util"
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
    fmt.Println(userName, newpassword)
    return true, nil
}