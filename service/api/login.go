package api

import (
    "github.com/astaxie/beego/validation"
    "GoModDemo/consts"
    "GoModDemo/util"
    "GoModDemo/service/authentication"
    "github.com/gin-gonic/gin"
    "net/http"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}


func Login(c *gin.Context) {
    appG := util.Gin{C: c}
    valid := validation.Validation{}
    username := c.Query("username")
    password := c.Query("password")

    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)
    if !ok {
        appG.Response(http.StatusOK, consts.INVALID_PARAMS, nil)
        return
    }

    authService := authentication.Auth{Username: username, Password: password}
    isExist, err := authService.Check()
    if err != nil {
        appG.Response(http.StatusOK, consts.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
        return
    }

    if !isExist {
        appG.Response(http.StatusOK, consts.ERROR_AUTH, nil)
        return
    }

    token, err := util.GenerateToken(username, password)
    if err != nil {
        appG.Response(http.StatusOK, consts.ERROR_AUTH_TOKEN, nil)
        return
    }

    appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
        "token": token,
    })
}