package api
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

// @Summary Login
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"Code":200,"Data":{},"Message":"ok"}"
// @Failure 200 {string} json "{"Code":,"Data":{},"Message":"ok"}"
// @Router /api/helper [get]
func IndexApi(c *gin.Context) {
	var t time.Time
	t = time.Now()
	str  := "你好!当前时间："+t.Format("2006-01-02 15:04:05")
	// c.String(http.StatusOK, str)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code": 200,
		"message": "你好！我是测试服务",
		"data": str,
	})
}
