package api

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/bill"
	"GoModDemo/util"
	"net/http"
	// "time"
	// "github.com/google/uuid"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"

)
// GetMenuByID 根据菜单ID获取菜单信息
// @Summary 根据菜单ID获取菜单信息
// @Tags Menu
// @Description 根据菜单ID获取菜单信息 请求主体: base64(ID=aaaa) 成功输出Menu
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ID formData string true "ID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{Menu},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetMenuByID [post]
func GetMenuByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据菜单ID获取菜单信息错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	ID := dists["ID"][0]
	if ID == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	temp, err := bill.GetMenuByID(ID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*temp)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s := string(b)
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetAllMenuInfo 前台条件获取菜单信息
// @Summary 前台条件获取菜单信息
// @Tags Menu
// @Description 前台条件获取菜单信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]Menu
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]Menu},"Message":""} or {"Code":-1,"Data":{},"Menu":"错误提示"}"
// @Router  /GetAllMenuInfo [post]
func GetAllMenuInfo(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	ParameterStr := dists["ParameterStr"][0]
	PageSize := dists["PageSize"][0]
	CurrentPage := dists["CurrentPage"][0]
	pageSize, err := strconv.Atoi(PageSize)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "参数异常", nil)
		return
	}
	currentPage, err := strconv.Atoi(CurrentPage)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "参数异常", nil)
		return
	}
	data,num, err := bill.GetAllMenuInfo(ParameterStr, pageSize, currentPage)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*data)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s := string(b)
	appG.Response1(http.StatusOK, consts.SUCCESS, "", s,num)
}

// GetAllMenuViewInfo 前台条件获取菜单信息
// @Summary 前台条件获取菜单信息
// @Tags Menu
// @Description 前台条件获取菜单信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]Menu
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]MenuView},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllMenuViewInfo [post]
func GetAllMenuViewInfo(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	ParameterStr := dists["ParameterStr"][0]
	PageSize := dists["PageSize"][0]
	CurrentPage := dists["CurrentPage"][0]
	pageSize, err := strconv.Atoi(PageSize)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "参数异常", nil)
		return
	}
	currentPage, err := strconv.Atoi(CurrentPage)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "参数异常", nil)
		return
	}
	data,num, err := bill.GetAllMenuViewInfo(ParameterStr, pageSize, currentPage)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*data)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s := string(b)
	appG.Response1(http.StatusOK, consts.SUCCESS, "", s,num)
}


// GetCascaderMenu 根据用户ID获取用户菜单信息
// @Summary 根据用户ID获取用户菜单信息
// @Tags User
// @Description 根据用户ID获取用户菜单信息 请求主体: Null  成功输出[]MenuTree
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]MenuTree},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetCascaderMenu [post]
func GetCascaderMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg:=""
	s:=""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	
	isOk,err:= util.RedisExists("GetCascaderMenu")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s,err=util.GetRedisString("GetCascaderMenu")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetCascaderMenu()
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}
		b, err := json.Marshal(*data)
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}
		s = string(b)
		err= util.SetRedisAnyEx("GetCascaderMenu",s,"180")

		if err != nil {
			errMsg=err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}
// GetMenuAllCount 获取菜单总条数
// @Summary 获取菜单总条数
// @Tags User
// @Description 获取菜单总条数 请求主体: Null  成功输出 int
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{int},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetMenuAllCount [post]
func GetMenuAllCount(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg:=""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
		data, err := bill.GetMenuAllCount()
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}
		
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}	
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, data)
}

// DeleteMenu 删除菜单信息
// @Summary 删除菜单信息
// @Tags Menu
// @Description 删除菜单信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteMenu [post]
func DeleteMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	str := dists["str"][0]
	if str==""{
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	date   := strings.Split(str, ",")
	err = bill.DeleteMenu(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// AddMenu 新增菜单信息
// @Summary 新增菜单信息
// @Tags Menu
// @Description 新增菜单信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddMenu [post]
func AddMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	str := dists["str"][0]
	if str==""{
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	var date  model.Menu
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return 
	}
	err = bill.AddMenu(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateMenu 修改菜单信息
// @Summary 修改菜单信息
// @Tags Menu
// @Description 修改菜单信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateMenu [post]
func UpdateMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	str := dists["str"][0]
	if str==""{
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	var date  model.Menu
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return 
	}
	err = bill.UpdateMenu(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}
