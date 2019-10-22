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
	// "strconv"
	// "strings"
)

// GetAllDeptForTransfer  获取所有部门清单
// @Summary 获取所有部门清单
// @Tags Dept
// @Description 获取所有部门清单 请求主体: Null  成功输出[]TransferModel
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":{[]TransferModel},"Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router  /GetAllDeptForTransfer [post]
func GetAllDeptForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	errMsg := ""
	s := ""
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "错误", nil)
		}
	}()
	isOk, err := util.RedisExists("GetAllDeptForTransfer")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	if isOk {
		s, err = util.GetRedisString("GetAllDeptForTransfer")
		if err != nil {
			appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
			return
		}

	} else {
		data, err := bill.GetAllDeptForTransfer()
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
		err = util.SetRedisAnyEx("GetAllDeptForTransfer", s, "180")

		if err != nil {
			errMsg = err.Error()
		}

	}
	appG.Response(http.StatusOK, consts.SUCCESS, errMsg, s)
}

// GetDeptByUserIdForTransfer 根据用户ID获取用户部门信息
// @Summary 根据用户ID获取用户部门信息
// @Tags Dept
// @Description 根据用户ID获取用户部门信息 请求主体: base64(userId=aaaa) 成功输出字符串
// @Accept mpfd
// @Param Token formData string true "Token"
// @Produce  json
// @Success 200 {string} json "{"Code":1,"Data":"aaa,bbb,ccc","Message":""} or {"Code":-1,"Data":{},"Message":"错误提示"}"
// @Router /GetDeptByUserIdForTransfer [post]
func GetDeptByUserIdForTransfer(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据用户ID获取用户部门信息", nil)
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
	s, err := bill.GetDeptByUserIdForTransfer(userId)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
		return
	}
	appG.Response(http.StatusOK, consts.SUCCESS, "", s)

}
