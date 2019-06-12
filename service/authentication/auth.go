package authentication

import "fmt"

type Auth struct {
    Username string
    Password string
}

func (a *Auth) Check() (bool, error) {
    userName := a.Username
    passWord := a.Password
  // todo：实现自己的鉴权逻辑
    fmt.Println(userName, passWord)
    return true, nil
}