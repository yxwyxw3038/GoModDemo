package api

import (
	"GoModDemo/consts"
	"GoModDemo/model"
	"GoModDemo/util"
	"GoModDemo/bill"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
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
		// user.AccountName="0000" 
		// user.PassWord="A390113400778BE79BC2B853A4E808CA"  
		// user.RealName="0000"
		// user.CreateBy="admin"
		// user.CreateTime=t.Format("2006-01-02 15:04:05")
		// user.UpdateBy="admin"
		// user.UpdateTime=t.Format("2006-01-02 15:04:05")
		// user.IsAble=1
		// user.IfChangePwd=1
		// _, err =db.Insert(&user)

		
		if err != nil {
			str = err.Error()
		} else {

			str = "你好!当前时间：" + t.Format("2006-01-02 15:04:05") 
		}
		// _,err=bill.GetUserMenu("f86bec36-18f9-4444-ab48-d645383956fd")
		// if err != nil {
		//   fmt.Println(err)
		// }
		list := make([]model.FilterModel, 0)
		 var temp1 model.FilterModel 
		 temp1.Column="UpdateTime"
		 temp1.Value="2019-09-26 16:47:58"
		 temp1.Action=">="
		 temp1.Logic="AND"
		 temp1.DataType="D"
		 list=append(list,temp1)
		 var temp2 model.FilterModel 
		 temp2.Column="AccountName"
		 temp2.Value="0000"
		 temp2.Action="="
		 temp2.Logic="AND"
		 temp2.DataType="S"
		 list=append(list,temp2)
		 b, err := json.Marshal(list)
		 if err != nil {
			fmt.Println(err)
			return
		}
		s:= string(b)
         _,err=bill.GetAllUserInfo(s,1,1)
		if err != nil {
		  fmt.Println(err)
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
