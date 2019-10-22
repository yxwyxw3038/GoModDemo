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

// GetAllRoleForTransfer  获取所有权限清单
// @Summary 获取所有权限清单
// @Tags Role
// @Description 获取所有权限清单 请求主体: Null  成功输出[]TransferModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TransferModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllRoleForTransfer [post]
func GetAllRoleForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	isOk, err := util.RedisExists("GetAllRoleForTransfer")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetAllRoleForTransfer")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetAllRoleForTransfer()
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
		err = util.SetRedisAnyEx("GetAllRoleForTransfer", s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetRoleByUserIdForTransfer 根据用户ID获取用户权限信息
// @Summary 根据用户ID获取用户权限信息
// @Tags Role
// @Description 根据用户ID获取用户权限信息 请求主体: base64(userId=aaaa) 成功输出字符串
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"aaa,bbb,ccc","Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetDeptByUserIdForTransfer [post]
func GetRoleByUserIdForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据用户ID获取用户权限信息", nil)
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
	s, err := bill.GetRoleByUserIdForTransfer(userId)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// GetAllRoleInfo 前台条件获取权限信息
// @Summary 前台条件获取权限信息
// @Tags Role
// @Description 前台条件获取权限信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]User},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllRoleInfo [post]
func GetAllRoleInfo(c *gin.Context) {
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

	data, err := bill.GetAllRoleInfo(ParameterStr, pageSize, currentPage)
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

// DeleteRole 删除权限信息
// @Summary 删除权限信息
// @Tags Role
// @Description 删除权限信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteRole [post]
func DeleteRole(c *gin.Context) {
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
	err = bill.DeleteRole(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// AddRole 新增权限信息
// @Summary 新增权限信息
// @Tags Role
// @Description 新增权限信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddRole [post]
func AddRole(c *gin.Context) {
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
	var date model.Role
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.AddRole(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateRole 修改权限信息
// @Summary 修改权限信息
// @Tags Role
// @Description 修改权限信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateRole [post]
func UpdateRole(c *gin.Context) {
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
	var date model.Role
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.UpdateRole(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// GetRoleByID 根据权限ID获取权限信息
// @Summary 根据权限ID获取权限信息
// @Tags Role
// @Description 根据权限ID获取权限信息 请求主体: base64(ID=aaaa) 成功输出Role
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ID formData string true "ID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{Role},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetRoleByID [post]
func GetRoleByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据权限ID获取权限信息错误", nil)
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
	role, err := bill.GetRoleByID(ID)
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

// GetMenuByRoleIdForTree 根据权限ID获取菜单树所有根结点
// @Summary 根据权限ID获取菜单树所有根结点
// @Tags Role
// @Description 根据权限ID获取菜单树所有根结点 请求主体: base64(userId=aaaa) 成功输出字符串
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"aaa,bbb,ccc","Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetMenuByRoleIdForTree [post]
func GetMenuByRoleIdForTree(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据权限ID获取菜单树所有根结点", nil)
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
	s, err := bill.GetMenuByRoleIdForTree(ID)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}

// SetMenuRole 设置权限菜单
// @Summary 设置权限菜单
// @Tags Role
// @Description 设置权限菜单 请求主体: base64(roleId=aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param roleId formData string true "权限ID""
// @Param menuStr formData string true "菜单ID，拼着接字符串"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /SetMenuRole [post]
func SetMenuRole(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "设置权限菜单", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	roleId := dists["roleId"][0]
	menuStr := dists["menuStr"][0]
	if roleId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	err = bill.SetMenuRole(roleId, menuStr)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, "", "")

}

// SetButtonByMenuIdRoleId 根据菜单ID和角色ID设置按钮权限
// @Summary 根据菜单ID和角色ID设置按钮权限
// @Tags Role
// @Description 根据菜单ID和角色ID设置按钮权限 请求主体: base64(menuId=aaaa&menuId=aaaa) 成功输出User
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param menuId formData string true "菜单ID"
// @Param roleId formData string true "权限ID"
// @Param menuStr formData string true "按钮ID，拼着接字符串"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /SetButtonByMenuIdRoleId [post]
func SetButtonByMenuIdRoleId(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据菜单ID和角色ID设置按钮权限", nil)
		}
	}()
	dists, err := appG.ParseQuery()
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, "数据包解密失败", nil)
		return
	}
	menuId := dists["menuId"][0]
	roleId := dists["roleId"][0]
	buttonStr := dists["buttonStr"][0]
	if roleId == "" {
		appG.Response(http.StatusOK, consts.ERROR, "参数为空", nil)
		return
	}
	err = bill.SetButtonByMenuIdRoleId(menuId, roleId, buttonStr)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}

	appG.Response(http.StatusOK, consts.SUCCESS, "", "")

}
