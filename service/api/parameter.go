package api

import (
	"GoModDemo/bill"
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"net/http"
	"strconv"
	"strings"

	// "time"
	// "github.com/google/uuid"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// GetTreeParameter 获取参数树信息
// @Summary 获取参数树信息
// @Tags Parameter
// @Description 获取参数树信息 请求主体: Null  成功输出[]TreeModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TreeModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetTreeParameter [post]
func GetTreeParameter(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()

	isOk, err := util.RedisExists("GetTreeParameter")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetTreeParameter")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetTreeParameter()
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
		err = util.SetRedisAnyEx("GetTreeParameter", s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetAllParameterInfo 前台条件获取参数信息
// @Summary 前台条件获取参数信息
// @Tags Parameter
// @Description 前台条件获取参数信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]Parameter
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]Parameter},"Message":""} or {"Code":-1,"Data":{},"Parameter":"错误提示"}"
// @Router  /GetAllParameterInfo [post]
func GetAllParameterInfo(c *gin.Context) {
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
	data, num, err := bill.GetAllParameterInfo(ParameterStr, pageSize, currentPage)
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
	appG.Response1(http.StatusOK, consts.SUCCESS, "", s, num)
}

// GetParameterByID 根据参数ID获取参数信息
// @Summary 根据参数ID获取参数信息
// @Tags Parameter
// @Description 根据参数ID获取参数信息 请求主体: base64(ID=aaaa) 成功输出Parameter
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ID formData string true "ID"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{Parameter},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetParameterByID [post]
func GetParameterByID(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据参数ID获取参数信息", nil)
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
	temp, err := bill.GetParameterByID(ID)
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

// DeleteParameter 删除参数信息
// @Summary 删除参数信息
// @Tags Parameter
// @Description 删除参数信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteParameter [post]
func DeleteParameter(c *gin.Context) {
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
	err = bill.DeleteParameter(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// AddParameter 新增参数信息
// @Summary 新增参数信息
// @Tags Parameter
// @Description 新增参数信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddParameter [post]
func AddParameter(c *gin.Context) {
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
	var date model.Parameter
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.AddParameter(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateParameter 修改参数信息
// @Summary 修改参数信息
// @Tags Parameter
// @Description 修改参数信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateParameter [post]
func UpdateParameter(c *gin.Context) {
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
	var date model.Parameter
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.UpdateParameter(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}
