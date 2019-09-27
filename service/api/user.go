package api

import (
	"GoModDemo/consts"
	// "GoModDemo/model"
	"GoModDemo/util"
	"net/http"
	"GoModDemo/bill"
	// "time"
	// "github.com/google/uuid"
	"encoding/json"
	"github.com/gin-gonic/gin"
)
// GetUserByID 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Tags User
// @Description 根据用户ID获取用户信息 请求主体: ID=base64(aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token" 
// @Param ID formData string true "ID" 
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{User},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetUserByID [post]
func GetUserByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func(){
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据用户ID获取用户信息错误", nil)
		}
	}()
    dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	ID:=dists["ID"][0]
	if ID=="" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
    user, err :=bill.GetUserByID(ID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*user)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s:= string(b)
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}
// GetUserMenu 根据用户ID获取用户菜单信息
// @Summary 根据用户ID获取用户菜单信息
// @Tags User
// @Description 根据用户ID获取用户菜单信息 请求主体: userID=base64(aaaa) 成功输出[]MenuTree
// @Accept mpfd
// @Param Token formData string true "Token" 
// @Param userID formData string true "userID" 
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]MenuTree},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetUserMenu [post]
func GetUserMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func(){
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
    dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	userID:=dists["userID"][0]
	if userID=="" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
    data, err :=bill.GetUserMenu(userID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*data)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s:= string(b)
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}