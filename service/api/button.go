package api

import (
	"GoModDemo/consts"
	// "GoModDemo/model"
	"GoModDemo/bill"
	"GoModDemo/util"
	"net/http"

	// "time"
	// "github.com/google/uuid"
	"encoding/json"

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
// @Tags Dept
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
// @Tags Dept
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
// @Tags Dept
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
// @Tags Dept
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
