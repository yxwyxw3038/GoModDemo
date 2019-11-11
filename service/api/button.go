package api

import (
	"GoModDemo/consts"
	"GoModDemo/model"

	// "GoModDemo/model"
	"GoModDemo/bill"
	"GoModDemo/util"
	"net/http"
	"strings"

	// "time"
	// "github.com/google/uuid"
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetButtonByMenuIdAndUserId 根据用户ID和菜单ID获取相应权限的按钮信息
// @Summary 根据用户ID和菜单ID获取相应权限的按钮信息
// @Tags Button
// @Description 根据用户ID获取用户菜单信息 请求主体:  base64(userId=aaaa&menuId=aaaa)  成功输出[]Button
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param userId formData string true "userId"
// @Param menuId formData string true "menuId"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]Button},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetButtonByMenuIdAndUserId [post]
func GetButtonByMenuIdAndUserId(c *gin.Context) {
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
	userId := dists["userId"][0]
	if userId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	menuId := dists["menuId"][0]
	if userId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	key := "GetButtonByMenuIdAndUserId" + userId + menuId
	isOk, err := util.RedisExists(key)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString(key)
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetButtonByMenuIdAndUserId(menuId, userId)
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
		err = util.SetRedisAnyEx(key, s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetAllButtonForTransfer  获取所有菜单清单
// @Summary 获取所有菜单清单
// @Tags Button
// @Description 获取所有菜单清单 请求主体: Null  成功输出[]TransferModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TransferModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllButtonForTransfer [post]
func GetAllButtonForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	isOk, err := util.RedisExists("GetAllButtonForTransfer")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetAllButtonForTransfer")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetAllButtonForTransfer()
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
		err = util.SetRedisAnyEx("GetAllButtonForTransfer", s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetButtonByMenuIdnForTransfer 根据菜单ID获取按钮信息
// @Summary 根据菜单ID获取按钮信息
// @Tags Button
// @Description 根据菜单ID获取按钮信息 请求主体: base64(menuId=aaaa) 成功输出字符串
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param menuId formData string true "menuId"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"aaa,bbb,ccc","Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetButtonByMenuIdnForTransfer [post]
func GetButtonByMenuIdnForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据菜单ID获取按钮信息", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	menuId := dists["menuId"][0]
	if menuId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	s, err := bill.GetButtonByMenuIdnForTransfer(menuId)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetAllButtonByMenuIdForTransfer  根据菜单ID获取按钮清单
// @Summary 根据菜单ID获取按钮清单
// @Tags Button
// @Description 根据菜单ID获取按钮清单 请求主体: Null  成功输出[]TransferModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TransferModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllButtonByMenuIdForTransfer [post]
func GetAllButtonByMenuIdForTransfer(c *gin.Context) {

	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据菜单ID获取按钮信息", nil)
		}
	}()
	s := ""
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}

	menuId := dists["menuId"][0]
	if menuId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	data, err := bill.GetAllButtonByMenuIdForTransfer(menuId)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if data != nil {
		b, err := json.Marshal(*data)
		s = string(b)
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}
	}

	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetButtonByMenuIdRoleIdForTransfer 根据菜单ID和权限ID获取按钮信息
// @Summary  根据菜单ID和权限ID获取按钮信息
// @Tags Button
// @Description  根据菜单ID和权限ID获取按钮信息 请求主体: base64(menuId=aaaa&roleId=bbbb) 成功输出字符串
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param menuId formData string true "menuId"
// @Param roleId formData string true "roleId"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"aaa,bbb,ccc","Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetButtonByMenuIdRoleIdForTransfer [post]
func GetButtonByMenuIdRoleIdForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, " 根据菜单ID和权限ID获取按钮信息", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	menuId := dists["menuId"][0]
	roleId := dists["roleId"][0]
	if menuId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	s, err := bill.GetButtonByMenuIdRoleIdForTransfer(menuId, roleId)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetAllButtonInfo 前台条件获取按钮信息
// @Summary 前台条件获取按钮信息
// @Tags Button
// @Description 前台条件获取按钮信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]Button
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]Button},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllButtonInfo [post]
func GetAllButtonInfo(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "前台条件获取按钮信息错误", nil)
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

	data, err := bill.GetAllButtonInfo(ParameterStr, pageSize, currentPage)
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

// DeleteButton 删除按钮信息
// @Summary 删除按钮信息
// @Tags Button
// @Description 删除按钮信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteButton [post]
func DeleteButton(c *gin.Context) {
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
	err = bill.DeleteButton(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// AddButton 新增按钮信息
// @Summary 新增按钮信息
// @Tags Button
// @Description 新增按钮信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddButton [post]
func AddButton(c *gin.Context) {
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
	var date model.Button
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.AddButton(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateButton 修改按钮信息
// @Summary 修改按钮信息
// @Tags Button
// @Description 修改按钮信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateButton [post]
func UpdateButton(c *gin.Context) {
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
	var date model.Button
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.UpdateButton(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// GetButtonByID 根据权限ID获取按钮信息
// @Summary 根据权限ID获取按钮信息
// @Tags Button
// @Description 根据权限ID获取按钮信息 请求主体: base64(ID=aaaa) 成功输出Button
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ID formData string true "ID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{Button},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetButtonByID [post]
func GetButtonByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据权限ID获取按钮信息错误", nil)
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
	role, err := bill.GetButtonByID(ID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	b, err := json.Marshal(*role)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	s := string(b)
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetButtonAllCount 获取按钮总条数
// @Summary 获取按钮总条数
// @Tags Button
// @Description 获取按钮总条数 请求主体: Null  成功输出 int
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{int},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetButtonAllCount [post]
func GetButtonAllCount(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	data, err := bill.GetButtonAllCount()
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
