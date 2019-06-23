package api

import (
	"GoModDemo/consts"
	"GoModDemo/service/authentication"
	"GoModDemo/util"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// Login 登录验证服务
// @Summary 登录验证服务
// @Tags User
// @Description 验证用户名密码有效性 accountName=aaa&passWord=base64(bbb) 成功输出Token
// @Accept mpfd
// @Param accountName formData string true "用户名" default(yxw)
// @Param passWord formData string true "密码" default(123)
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{Token},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /Login [post]
func Login(c *gin.Context) {
	logger := util.InitZapLog()
	logger.Debug("开始登录验证！")
	appG := util.Gin{C: c}
	valid := validation.Validation{}
	str, err := appG.GetBase64Body()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "解密失败", nil)
		return
	}
	fmt.Printf(str)
	dists, err := url.ParseQuery(str)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "反URL化失败", nil)
		return
	}
	username := dists["accountName"][0]
	password := dists["passWord"][0]
	unsz, err := base64.StdEncoding.DecodeString(password)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "密码解密异常", nil)
		return
	}
	unpassword := string(unsz)
	a := auth{Username: username, Password: unpassword}
	ok, _ := valid.Valid(&a)
	if !ok {
		appG.Response(http.StatusOK, consts.ERROR, "JWT验证失败", nil)
		return
	}

	authService := authentication.Auth{Username: username, Password: unpassword}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "JWT验证失败", nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusOK, consts.ERROR, "JWT验证失败", nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "JWT验证失败", nil)
		return
	}

	// appG.Response(http.StatusOK, consts.SUCCESS, map[string]string{
	//     "token": token,
	// })
	appG.Response(http.StatusOK, consts.SUCCESS, "JWT验证通过", token)
}
