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

// GetAllNoticeInfo 前台条件获取通知单信息
// @Summary 前台条件获取通知单信息
// @Tags Notice
// @Description 前台条件获取通知单信息 请求主体: base64(ParameterStr=aaaa&PageSize=10&CurrentPage=1) 成功输出[]NoticeViewModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param ParameterStr formData string true "ParameterStr"
// @Param PageSize formData int true "PageSize"
// @Param CurrentPage formData int true "CurrentPage"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]NoticeViewModel},"Message":""} or {"Code":-1,"Data":{},"Parameter":"错误提示"}"
// @Router  /GetAllNoticeInfo [post]
func GetAllNoticeInfo(c *gin.Context) {
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
	data, num, err := bill.GetAllNoticeInfo(ParameterStr, pageSize, currentPage)
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

// AddNotice 新增通知单信息
// @Summary 新增通知单信息
// @Tags Notice
// @Description 新增通知单信息 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /AddNotice [post]
func AddNotice(c *gin.Context) {
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
	var date model.NoticeBillModel
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.AddNotice(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// UpdateNotice 修改通知单
// @Summary 修改通知单
// @Tags Notice
// @Description 修改通知单 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /UpdateNotice [post]
func UpdateNotice(c *gin.Context) {
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
	var date model.NoticeBillModel
	err = json.Unmarshal([]byte(str), &date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	err = bill.UpdateNotice(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}

// DeleteNotice 删除通知单
// @Summary 删除通知单
// @Tags Notice
// @Description 删除通知单 请求主体: base64(str=aaaa) 成功输出null
// @Accept mpfd
// @Param Token formData string true "Token"
// @Param str formData string true "str"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /DeleteNotice [post]
func DeleteNotice(c *gin.Context) {
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
	err = bill.DeleteNotice(date)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", "")
}
