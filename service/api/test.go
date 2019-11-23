package api

import (
	"GoModDemo/consts"
	"GoModDemo/util"
	"net/http"
	"time"

	// "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

// }
// @Summary Login
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"Code":200,"Data":{},"Message":"ok"}"
// @Failure 200 {string} json "{"Code":,"Data":{},"Message":"ok"}"
// @Router /api/helper [get]
func IndexApi(c *gin.Context) {
	appG := util.Gin{C: c}
	defer func() {
		if p := recover(); p != nil {
			appG.Response(http.StatusOK, consts.ERROR, "根据用户ID获取用户权限信息", nil)
		}
	}()
	var t time.Time
	t = time.Now()
	var str string
	str0, err := util.AesEncryptStr("你好！")
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
	}
	str1, err := util.AesDecryptStr(str0)
	if err != nil {
		appG.Response(http.StatusOK, consts.ERROR, err.Error(), nil)
	}
	str = "你好!当前时间：" + t.Format("2006-01-02 15:04:05")
	str += "  密文：" + str0
	str += "  明文：" + str1
	appG.Response(http.StatusOK, consts.SUCCESS, "获取时间成功", str)
}
