package api

import (
	"GoModDemo/consts"
	// "GoModDemo/model"
	"GoModDemo/util"
	"net/http"
	"time"
	// "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

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
	db, err := util.OpenDB()
	if err != nil {
		str = err.Error()
	} else {
		_, err := db.Table("User").Count()
		// newUuid := uuid.New()
		// newUuidStr := newUuid.String()
		// var user model.User
		// user.ID=newUuidStr
		// user.AccountName="admin" 
		// user.PassWord="F6FDFFE48C908DEB0F4C3BD36C032E72"  
		// user.RealName="系统管理员"
		// user.CreateBy="admin"
		// user.CreateTime=t.Format("2006-01-02 15:04:05")
		// user.UpdateBy="admin"
		// user.UpdateTime=t.Format("2006-01-02 15:04:05")
		// user.IsAble=1
		// user.IfChangePwd=1
		// _, err :=db.Insert(&user)
		if err != nil {
			str = err.Error()
		} else {

			str = "你好!当前时间：" + t.Format("2006-01-02 15:04:05") 
		}
	}
	// c.String(http.StatusOK, str)
	// c.JSON(http.StatusOK, gin.H{
	// 	"success": true,
	// 	"code": 1,
	// 	"message": "你好！我是测试服务",
	// 	"data": str,
	// })
	appG.Response(http.StatusOK, consts.SUCCESS, "获取时间成功", str)
}
