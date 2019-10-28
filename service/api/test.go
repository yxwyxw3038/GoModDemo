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
	var t time.Time
	t = time.Now()
	var str string
	str = "你好!当前时间：" + t.Format("2006-01-02 15:04:05")
	appG.Response(http.StatusOK, consts.SUCCESS, "获取时间成功", str)
}
