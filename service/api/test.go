package api
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
)

//IndexApi 测试服务
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
