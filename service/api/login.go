package api

import (
	"GoModDemo/consts"
	"GoModDemo/service/authentication"
	"GoModDemo/util"
	"fmt"
	"net/http"
	"net/url"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"encoding/base64"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func Login(c *gin.Context) {
	appG := util.Gin{C: c}
	valid := validation.Validation{}	
	str, err := appG.GetBase64Body()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, nil)
		return
	}
	fmt.Printf(str)
	dists,err:=	url.ParseQuery(str)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, nil)
		return
	}
	username := dists["accountName"][0]
	password := dists["passWord"][0]
	unsz, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "密码解密异常")
		return
	}
	unpassword:=string(unsz)
	a := auth{Username: username, Password: unpassword}
	ok, _ := valid.Valid(&a)
	if !ok {
		appG.Response(http.StatusOK, consts.INVALID_PARAMS, nil)
		return
	}

	authService := authentication.Auth{Username: username, Password: unpassword}
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

	// appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
	//     "token": token,
	// })
	appG.Response(http.StatusOK, consts.SUCCESS, token)
}
