package api

import (
	"GoModDemo/bill"
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"net/http"

	// "time"
	// "github.com/google/uuid"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetUserByID 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Tags User
// @Description 根据用户ID获取用户信息 请求主体: base64(ID=aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ID formData string true "ID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{User},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetUserByID [post]
func GetUserByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据用户ID获取用户信息错误", nil)
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
	user, err := bill.GetUserByID(ID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*user)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s := string(b)
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetUserMenu 根据用户ID获取用户菜单信息
// @Summary 根据用户ID获取用户菜单信息
// @Tags User
// @Description 根据用户ID获取用户菜单信息 请求主体:  base64(userID=aaaa)  成功输出[]MenuTree
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param userID formData string true "userID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]MenuTree},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetUserMenu [post]
func GetUserMenu(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
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
	userID := dists["userID"][0]
	if userID == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	isOk, err := util.RedisExists("GetUserMenu" + userID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetUserMenu" + userID)
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetUserMenu(userID)
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
		err = util.SetRedisAnyEx("GetUserMenu"+userID, s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetAllUserInfo 前台条件获取用户信息
// @Summary 前台条件获取用户信息
// @Tags User
// @Description 前台条件获取用户信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]User},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllUserInfo [post]
func GetAllUserInfo(c *gin.Context) {
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
	if ParameterStr == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}

	data, err := bill.GetAllUserInfo(ParameterStr, pageSize, currentPage)
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
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)
}

// GetAllUserViewInfo 前台条件获取用户信息
// @Summary 前台条件获取用户信息
// @Tags User
// @Description 前台条件获取用户信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]UserView},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllUserViewInfo [post]
func GetAllUserViewInfo(c *gin.Context) {
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
	if ParameterStr == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}

	data, err := bill.GetAllUserViewInfo(ParameterStr, pageSize, currentPage)
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
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)
}

// GetUserLogininfoByToken 验证Token有效性
// @Summary 验证Token有效性
// @Tags User
// @Description 验证Token有效性 请求主体:为空 成功输出 空
// @Accept mpfd
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"","Message":""}"
// @Router  /GetUserLogininfoByToken [post]
func GetUserLogininfoByToken(c *gin.Context) {

	appG := util.Gin{C: c}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")

}

// DeleteUser 删除用户信息
// @Summary 删除用户信息
// @Tags User
// @Description 删除用户信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteUser [post]
func DeleteUser(c *gin.Context) {
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
	if str == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	date := strings.Split(str, ",")
	err = bill.DeleteUser(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// AddUser 新增用户信息
// @Summary 新增用户信息
// @Tags User
// @Description 新增用户信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddUser [post]
func AddUser(c *gin.Context) {
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
	if str == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	var date model.User
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.AddUser(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateUser 修改用户信息
// @Summary 修改用户信息
// @Tags User
// @Description 修改用户信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateUser [post]
func UpdateUser(c *gin.Context) {
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
	if str == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	var date model.User
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.UpdateUser(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// SetUserDept 设置用户部门
// @Summary 设置用户部门
// @Tags User
// @Description 设置用户部门 请求主体: base64(ID=aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param userId formData string true "用户ID""
// @Param userId formData string true "部门ID，拼着接字符串"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /SetUserDept [post]
func SetUserDept(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "设置用户部门", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	userId := dists["userId"][0]
	deptStr := dists["deptStr"][0]
	if userId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	err = bill.SetUserDept(userId, deptStr)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, "", "")

}

// SetUserRole 设置用户权限
// @Summary 设置用户权限
// @Tags User
// @Description 设置设置用户权限用户部门 请求主体: base64(ID=aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param userId formData string true "用户ID""
// @Param userId formData string true "权限ID，拼着接字符串"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /SetUserRole [post]
func SetUserRole(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "设置用户权限", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	userId := dists["userId"][0]
	roleStr := dists["roleStr"][0]
	if userId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	err = bill.SetUserRole(userId, roleStr)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, "", "")

}

// GetTreeUser 用户信息树
// @Summary 用户信息树
// @Tags Menu
// @Description 用户信息树 请求主体: Null  成功输出[]TreeNodeModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TreeNodeModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetTreeUser [post]
func GetTreeUser(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()

	isOk, err := util.RedisExists("GetTreeUser")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetTreeUser")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetTreeUser()
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
		err = util.SetRedisAnyEx("GetTreeUser", s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}
